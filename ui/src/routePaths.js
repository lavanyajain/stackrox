/**
 * Application route paths constants.
 */

import { resourceTypes, standardEntityTypes } from 'constants/entityTypes';

export const mainPath = '/main';
export const loginPath = '/login';
export const licenseStartUpPath = `/license`;
export const authResponsePrefix = '/auth/response/';

export const dashboardPath = `${mainPath}/dashboard`;
export const networkPath = `${mainPath}/network`;
export const violationsPath = `${mainPath}/violations/:alertId?`;
export const integrationsPath = `${mainPath}/integrations`;
export const policiesPath = `${mainPath}/policies/:policyId?`;
export const riskPath = `${mainPath}/risk/:deploymentId?`;
export const imagesPath = `${mainPath}/images/:imageId?`;
export const secretsPath = `${mainPath}/secrets/:secretId?`;
export const apidocsPath = `${mainPath}/apidocs`;
export const accessControlPath = `${mainPath}/access`;
export const licensePath = `${mainPath}/license`;
export const systemConfigPath = `${mainPath}/systemconfig`;

/**
 *Compliance-related route paths
 */
export const resourceTypesToUrl = {
    [resourceTypes.NAMESPACE]: 'namespaces',
    [resourceTypes.CLUSTER]: 'clusters',
    [resourceTypes.NODE]: 'nodes',
    [resourceTypes.DEPLOYMENT]: 'deployments',
    [standardEntityTypes.CONTROL]: 'controls'
};

export const compliancePath = `${mainPath}/compliance`;
const resourceMatcher = `(${Object.values(resourceTypesToUrl).join('|')})`;

export const nestedCompliancePaths = {
    DASHBOARD: `${compliancePath}/`,
    LIST: `${compliancePath}/:entityType`,
    CONTROL: `${compliancePath}/:entityType(controls)/:controlId/:listEntityType${resourceMatcher}?`,
    CLUSTER: `${compliancePath}/:entityType(clusters)/:entityId/:listEntityType${resourceMatcher}?`,
    NAMESPACE: `${compliancePath}/:entityType(namespaces)/:entityId/:listEntityType${resourceMatcher}?`,
    DEPLOYMENT: `${compliancePath}/:entityType(deployments)/:entityId/:listEntityType${resourceMatcher}?`,
    NODE: `${compliancePath}/:entityType(nodes)/:entityId/:listEntityType${resourceMatcher}?`
};
