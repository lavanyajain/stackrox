/* eslint-disable @typescript-eslint/dot-notation */
/* eslint-disable import/prefer-default-export */
import { AuthProvider, AuthProviderConfig, Group } from 'services/AuthService';

export type DisplayedAuthProvider = AuthProvider & {
    do_not_use_client_secret?: boolean;
    defaultRole?: string;
    groups?: Group[];
};

export function transformInitialValues(
    initialValues: DisplayedAuthProvider
): DisplayedAuthProvider {
    // TODO-ivan: eventually logic for different auth provider type should live
    // with the form component that renders form for the corresponding auth provider
    // type, probably makes sense to refactor after moving away from redux-form
    if (initialValues.type === 'oidc') {
        const alteredConfig = { ...initialValues.config };

        // backend doesn't return the exact value for the client secret for the security reasons,
        // instead it'll return some obfuscated data, but not an empty one
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        alteredConfig.clientOnly = {
            clientSecretStored: !!alteredConfig.client_secret,
        };

        if (initialValues.name) {
            // if it's an existing auth provider, then we're using the secret if we have it
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            alteredConfig.do_not_use_client_secret = !alteredConfig.client_secret;
        }

        // clean-up obfuscated value if any as we don't need to show it
        alteredConfig.client_secret = '';

        return {
            ...initialValues,
            config: alteredConfig,
        };
    }
    if (initialValues.type === 'saml') {
        const alteredConfig = { ...initialValues.config };
        // unless static config values are present, assume dynamic configuration is selected
        alteredConfig.configurationType = alteredConfig.idp_issuer ? 'static' : 'dynamic';
        return {
            ...initialValues,
            config: alteredConfig,
        };
    }
    return initialValues;
}

function populateDefaultValues(authProvider: AuthProvider): AuthProvider {
    const newInitialValues: DisplayedAuthProvider = { ...authProvider };
    newInitialValues.uiEndpoint = window.location.host;
    newInitialValues.enabled = true;
    if (authProvider.type === 'oidc') {
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        newInitialValues.config = { mode: 'auto', do_not_use_client_secret: false };
    }
    newInitialValues.groups = Array.isArray(authProvider.groups) ? [...authProvider.groups] : [];
    newInitialValues.groups.push({
        roleName: '',
        props: { authProviderId: '', key: '', value: '' },
    });

    return newInitialValues;
}

export function getInitialAuthProviderValues(authProvider: AuthProvider): DisplayedAuthProvider {
    const initialValues: DisplayedAuthProvider =
        !authProvider.name && !authProvider.id
            ? populateDefaultValues(authProvider)
            : { ...authProvider };

    const modifiedInitialValues = {
        ...transformInitialValues(initialValues),
    };

    return modifiedInitialValues;
}

export function transformValuesBeforeSaving(
    values: Record<string, string | string[] | boolean | AuthProviderConfig | Group[] | undefined>
): Record<string, string | string[] | boolean | AuthProviderConfig | Group[] | undefined> {
    if (values.type === 'oidc') {
        const alteredConfig = { ...(values.config as AuthProviderConfig) };

        // if client secret is stored on the backend and user didn't enter any value,
        // it means that user wants to preserve the stored secret, delete then
        const preserveStoredClientSecret =
            typeof alteredConfig.clientOnly === 'object' &&
            'clientSecretStored' in alteredConfig.clientOnly &&
            typeof alteredConfig.clientOnly.clientSecretStored === 'boolean' &&
            alteredConfig.clientOnly?.clientSecretStored &&
            !alteredConfig.client_secret;
        if (alteredConfig.do_not_use_client_secret || preserveStoredClientSecret) {
            delete alteredConfig.client_secret;
        }

        // backend expects only string values for the config
        alteredConfig.do_not_use_client_secret = alteredConfig.do_not_use_client_secret
            ? 'true'
            : 'false';

        // finally delete client only values
        delete alteredConfig.clientOnly;

        return {
            ...values,
            config: alteredConfig,
        };
    }
    if (values.type === 'saml') {
        const alteredConfig = { ...(values.config as AuthProviderConfig) };
        if (alteredConfig.configurationType === 'dynamic') {
            ['idp_issuer', 'idp_sso_url', 'idp_nameid_format', 'idp_cert_pem'].forEach(
                (p) => delete alteredConfig[p]
            );
        } else if (alteredConfig.configurationType === 'static') {
            delete alteredConfig.idp_metadata_url;
        }
        delete alteredConfig.configurationType; // that was UI only field

        return {
            ...values,
            config: alteredConfig,
        };
    }
    return values;
}

export function getGroupsByAuthProviderId(groups: Group[], id: string): Group[] {
    const filteredGroups = groups.filter(
        (group) =>
            group.props &&
            group.props.authProviderId &&
            group.props.authProviderId === id &&
            group.props.key !== ''
    );
    return filteredGroups;
}

export function mergeGroupsWithAuthProviders(
    authProviders: AuthProvider[],
    groups: Group[]
): AuthProvider[] {
    const authProvidersWithGroupsDict = authProviders.reduce((obj, item) => {
        // reset rules on each calculation
        // eslint-disable-next-line no-param-reassign
        item.groups = [];

        // comma operator is much faster than spread in a reduce loop
        // eslint-disable-next-line prettier/prettier, no-return-assign, no-param-reassign, no-sequences, @typescript-eslint/dot-notation
            return (obj[item.id] = item), obj;
    }, {});

    if (authProviders.length) {
        groups.forEach((group) => {
            if (group.props.key !== '') {
                authProvidersWithGroupsDict[group.props.authProviderId].groups.push(group);
            } else {
                authProvidersWithGroupsDict[group.props.authProviderId].defaultRole =
                    group.roleName;
            }
        });
    }

    return Object.values(authProvidersWithGroupsDict);
}

export function getDefaultRoleByAuthProviderId(groups: Group[], id: string): string {
    let defaultRoleGroups = groups.filter(
        (group) =>
            group.props &&
            group.props.authProviderId &&
            group.props.authProviderId === id &&
            group.props.key === '' &&
            group.props.value === ''
    );
    if (defaultRoleGroups.length) {
        return defaultRoleGroups[0].roleName;
    }
    // if there is no default role specified for this auth provider then use the global default role
    defaultRoleGroups = groups.filter((group) => !group.props);
    if (defaultRoleGroups.length) {
        return defaultRoleGroups[0].roleName;
    }
    return 'Admin';
}