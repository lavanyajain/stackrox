import React, { ReactElement, useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import {
    Breadcrumb,
    BreadcrumbItem,
    TabTitleText,
    Tabs,
    Tab,
    Title,
    PageSection,
    Spinner,
} from '@patternfly/react-core';

import { violationsPFBasePath } from 'routePaths';
import { fetchAlert } from 'services/AlertsService';
import BreadcrumbItemLink from 'Components/BreadcrumbItemLink';
import { preFormatPolicyFields } from 'Containers/Policies/Wizard/Form/utils';
import ViolationsDetails from 'Containers/Violations/SidePanel/ViolationsDetails';
import EnforcementDetails from 'Containers/Violations/Enforcement/Details';
import PolicyDetails from 'Containers/Policies/Wizard/Details/PolicyDetails';
import DeploymentDetails from 'Containers/Risk/DeploymentDetails';
import { Alert } from '../types/violationTypes';
import ViolationNotFoundPage from '../ViolationNotFoundPage';

function ViolationDetailsPage(): ReactElement {
    const [activeTabKey, setActiveTabKey] = useState(0);
    const [alert, setAlert] = useState<Alert>();
    const [isFetchingSelectedAlert, setIsFetchingSelectedAlert] = useState(false);

    const { alertId } = useParams();

    function handleTabClick(_, tabIndex) {
        setActiveTabKey(tabIndex);
    }

    // Make updates to the fetching state, and selected alert.
    useEffect(() => {
        setIsFetchingSelectedAlert(true);
        fetchAlert(alertId).then(
            (result) => {
                setAlert(result);
                setIsFetchingSelectedAlert(false);
            },
            () => {
                setAlert(undefined);
                setIsFetchingSelectedAlert(false);
            }
        );
    }, [alertId, setAlert, setIsFetchingSelectedAlert]);

    if (!alert) {
        if (isFetchingSelectedAlert) {
            return <Spinner isSVG />;
        }
        return <ViolationNotFoundPage />;
    }

    const { policy, deployment, resource, commonEntityInfo } = alert;
    const title = policy.name || 'Unknown violation';
    const { name: entityName } = resource || deployment || {};
    const resourceType = resource?.resourceType || commonEntityInfo?.resourceType || 'deployment';

    return (
        <PageSection variant="light" isFilled id="violation-details">
            <Breadcrumb className="pf-u-mb-md">
                <BreadcrumbItemLink to={violationsPFBasePath}>Violations</BreadcrumbItemLink>
                <BreadcrumbItem isActive>{title}</BreadcrumbItem>
            </Breadcrumb>
            <Title headingLevel="h1">{title}</Title>
            <Title headingLevel="h2" className="pf-u-mb-md">{`in "${
                entityName as string
            }" ${resourceType}`}</Title>
            <Tabs mountOnEnter activeKey={activeTabKey} onSelect={handleTabClick}>
                <Tab eventKey={0} title={<TabTitleText>Violation</TabTitleText>}>
                    <ViolationsDetails
                        violationId={alert.id}
                        violations={alert.violations}
                        processViolation={alert.processViolation}
                        lifecycleStage={alert.lifecycleStage}
                    />
                </Tab>
                {alert?.enforcement && (
                    <Tab eventKey={1} title={<TabTitleText>Enforcement</TabTitleText>}>
                        <EnforcementDetails alert={alert} />
                    </Tab>
                )}
                {alert?.deployment && (
                    <Tab eventKey={2} title={<TabTitleText>Deployment</TabTitleText>}>
                        <DeploymentDetails deployment={alert.deployment} />
                    </Tab>
                )}
                <Tab eventKey={3} title={<TabTitleText>Policy</TabTitleText>}>
                    <PolicyDetails policy={preFormatPolicyFields(alert.policy)} />
                </Tab>
            </Tabs>
        </PageSection>
    );
}

export default ViolationDetailsPage;