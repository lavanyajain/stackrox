import gql from 'graphql-tag';

export const CLUSTER_LIST_FRAGMENT = gql`
    fragment clusterListFields on Cluster {
        id
        name
        # cves
        status {
            orchestratorMetadata {
                version
            }
        }
        # createdAt
        namespaceCount
        deploymentCount
        policyCount
        policyStatus {
            status
        }
        # latestViolation
        # risk
    }
`;

export const CVE_LIST_FRAGMENT = gql`
    fragment cveListFields on EmbeddedVulnerability {
        cve
        cvss
        scoreVersion
        # Env. Impact
        # Impact score
        summary
        fixedByVersion
        isFixable
        lastScanned
        # published
        deploymentCount
        imageCount
        componentCount
    }
`;

export const DEPLOYMENT_LIST_FRAGMENT = gql`
    fragment deploymentListFields on Deployment {
        id
        name
        vulnerabilities: vulns {
            cve
            cvss
            isFixable
        }
        deployAlerts {
            time
        }
        failingPolicyCount
        policyStatus
        clusterName
        clusterId
        namespace
        namespaceId
        serviceAccount
        serviceAccountID
        secretCount
        imageCount
        priority
    }
`;
