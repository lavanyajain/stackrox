import React, { ReactElement } from 'react';
import { Tooltip } from '@patternfly/react-core';
import {
    ExclamationTriangleIcon,
    CheckIcon,
    LongArrowAltDownIcon,
    LongArrowAltUpIcon,
    TimesIcon,
} from '@patternfly/react-icons';

import { EffectiveAccessScopeState } from 'services/RolesService';

const excludedColor = 'var(--pf-global--danger-color--100)';
const includedColor = 'var(--pf-global--success-color--100)';
const unknownColor = 'var(--pf-global--warning-color--100)';

const excludedIcon = <TimesIcon color={excludedColor} />;
const excludedCluster = (
    <Tooltip
        content={
            <div>
                Not included: cluster
                <br />
                nor any of its namespaces
            </div>
        }
        isContentLeftAligned
    >
        {excludedIcon}
    </Tooltip>
);
const excludedNamespace = <Tooltip content="Not included: namespace">{excludedIcon}</Tooltip>;

const includedIcon = <CheckIcon color={includedColor} />;
const includedCluster = (
    <Tooltip
        content={
            <div>
                Included: cluster
                <br />
                and therefore all of its namespaces
            </div>
        }
        isContentLeftAligned
    >
        <span>
            {includedIcon}
            <LongArrowAltDownIcon color={includedColor} style={{ transform: 'rotate(-45deg)' }} />
        </span>
    </Tooltip>
);
const includedNamespace = <Tooltip content="Included: namespace">{includedIcon}</Tooltip>;

const partialCluster = (
    <Tooltip
        content={
            <div>
                Conditionally included: cluster
                <br />
                because at least one of its namespaces
            </div>
        }
        isContentLeftAligned
    >
        <span>
            {includedIcon}
            <LongArrowAltUpIcon color={includedColor} style={{ transform: 'rotate(-45deg)' }} />
        </span>
    </Tooltip>
);

const unknownState = (
    <Tooltip content="Unknown">
        <ExclamationTriangleIcon color={unknownColor} />
    </Tooltip>
);

export type EffectiveAccessScopeStateProps = {
    state: EffectiveAccessScopeState;
    isCluster: boolean;
};

function EffectiveAccessScopeStateIcon({
    state,
    isCluster,
}: EffectiveAccessScopeStateProps): ReactElement {
    switch (state) {
        case 'EXCLUDED':
            return isCluster ? excludedCluster : excludedNamespace;

        case 'INCLUDED':
            return isCluster ? includedCluster : includedNamespace;

        case 'PARTIAL':
            return partialCluster;

        default:
            return unknownState;
    }
}

export default EffectiveAccessScopeStateIcon;