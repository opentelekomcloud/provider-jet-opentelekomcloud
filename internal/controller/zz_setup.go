/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/terraform"

	v1 "github.com/enrrou/otc-provider-jet/internal/controller/antiddos/v1"
	configurationv1 "github.com/enrrou/otc-provider-jet/internal/controller/as/configurationv1"
	groupv1 "github.com/enrrou/otc-provider-jet/internal/controller/as/groupv1"
	policyv1 "github.com/enrrou/otc-provider-jet/internal/controller/as/policyv1"
	policyv2 "github.com/enrrou/otc-provider-jet/internal/controller/as/policyv2"
	volumev2 "github.com/enrrou/otc-provider-jet/internal/controller/blockstorage/volumev2"
	policyv3 "github.com/enrrou/otc-provider-jet/internal/controller/cbr/policyv3"
	vaultv3 "github.com/enrrou/otc-provider-jet/internal/controller/cbr/vaultv3"
	addonv3 "github.com/enrrou/otc-provider-jet/internal/controller/cce/addonv3"
	clusterv3 "github.com/enrrou/otc-provider-jet/internal/controller/cce/clusterv3"
	nodepoolv3 "github.com/enrrou/otc-provider-jet/internal/controller/cce/nodepoolv3"
	nodev3 "github.com/enrrou/otc-provider-jet/internal/controller/cce/nodev3"
	alarmrule "github.com/enrrou/otc-provider-jet/internal/controller/ces/alarmrule"
	bmsserverv2 "github.com/enrrou/otc-provider-jet/internal/controller/compute/bmsserverv2"
	bmstagsv2 "github.com/enrrou/otc-provider-jet/internal/controller/compute/bmstagsv2"
	floatingipassociatev2 "github.com/enrrou/otc-provider-jet/internal/controller/compute/floatingipassociatev2"
	floatingipv2 "github.com/enrrou/otc-provider-jet/internal/controller/compute/floatingipv2"
	instancev2 "github.com/enrrou/otc-provider-jet/internal/controller/compute/instancev2"
	keypairv2 "github.com/enrrou/otc-provider-jet/internal/controller/compute/keypairv2"
	secgroupv2 "github.com/enrrou/otc-provider-jet/internal/controller/compute/secgroupv2"
	servergroupv2 "github.com/enrrou/otc-provider-jet/internal/controller/compute/servergroupv2"
	volumeattachv2 "github.com/enrrou/otc-provider-jet/internal/controller/compute/volumeattachv2"
	backuppolicyv1 "github.com/enrrou/otc-provider-jet/internal/controller/csbs/backuppolicyv1"
	backupv1 "github.com/enrrou/otc-provider-jet/internal/controller/csbs/backupv1"
	clusterv1 "github.com/enrrou/otc-provider-jet/internal/controller/css/clusterv1"
	snapshotconfigurationv1 "github.com/enrrou/otc-provider-jet/internal/controller/css/snapshotconfigurationv1"
	trackerv1 "github.com/enrrou/otc-provider-jet/internal/controller/cts/trackerv1"
	instancev1 "github.com/enrrou/otc-provider-jet/internal/controller/dcs/instancev1"
	instancev3 "github.com/enrrou/otc-provider-jet/internal/controller/dds/instancev3"
	hostv1 "github.com/enrrou/otc-provider-jet/internal/controller/deh/hostv1"
	groupv1dms "github.com/enrrou/otc-provider-jet/internal/controller/dms/groupv1"
	instancev1dms "github.com/enrrou/otc-provider-jet/internal/controller/dms/instancev1"
	queuev1 "github.com/enrrou/otc-provider-jet/internal/controller/dms/queuev1"
	ptrrecordv2 "github.com/enrrou/otc-provider-jet/internal/controller/dns/ptrrecordv2"
	recordsetv2 "github.com/enrrou/otc-provider-jet/internal/controller/dns/recordsetv2"
	zonev2 "github.com/enrrou/otc-provider-jet/internal/controller/dns/zonev2"
	instancev1ecs "github.com/enrrou/otc-provider-jet/internal/controller/ecs/instancev1"
	volumev3 "github.com/enrrou/otc-provider-jet/internal/controller/evs/volumev3"
	firewallgroupv2 "github.com/enrrou/otc-provider-jet/internal/controller/fw/firewallgroupv2"
	policyv2fw "github.com/enrrou/otc-provider-jet/internal/controller/fw/policyv2"
	rulev2 "github.com/enrrou/otc-provider-jet/internal/controller/fw/rulev2"
	agencyv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/agencyv3"
	credentialv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/credentialv3"
	groupmembershipv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/groupmembershipv3"
	groupv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/groupv3"
	mappingv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/mappingv3"
	projectv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/projectv3"
	protocolv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/protocolv3"
	providerv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/providerv3"
	roleassignmentv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/roleassignmentv3"
	rolev3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/rolev3"
	usergroupmembershipv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/usergroupmembershipv3"
	userv3 "github.com/enrrou/otc-provider-jet/internal/controller/identity/userv3"
	imageaccessacceptv2 "github.com/enrrou/otc-provider-jet/internal/controller/images/imageaccessacceptv2"
	imageaccessv2 "github.com/enrrou/otc-provider-jet/internal/controller/images/imageaccessv2"
	imagev2 "github.com/enrrou/otc-provider-jet/internal/controller/images/imagev2"
	dataimagev2 "github.com/enrrou/otc-provider-jet/internal/controller/ims/dataimagev2"
	imagev2ims "github.com/enrrou/otc-provider-jet/internal/controller/ims/imagev2"
	grantv1 "github.com/enrrou/otc-provider-jet/internal/controller/kms/grantv1"
	keyv1 "github.com/enrrou/otc-provider-jet/internal/controller/kms/keyv1"
	certificatev2 "github.com/enrrou/otc-provider-jet/internal/controller/lb/certificatev2"
	certificatev3 "github.com/enrrou/otc-provider-jet/internal/controller/lb/certificatev3"
	l7policyv2 "github.com/enrrou/otc-provider-jet/internal/controller/lb/l7policyv2"
	l7rulev2 "github.com/enrrou/otc-provider-jet/internal/controller/lb/l7rulev2"
	listenerv2 "github.com/enrrou/otc-provider-jet/internal/controller/lb/listenerv2"
	listenerv3 "github.com/enrrou/otc-provider-jet/internal/controller/lb/listenerv3"
	loadbalancerv2 "github.com/enrrou/otc-provider-jet/internal/controller/lb/loadbalancerv2"
	loadbalancerv3 "github.com/enrrou/otc-provider-jet/internal/controller/lb/loadbalancerv3"
	memberv2 "github.com/enrrou/otc-provider-jet/internal/controller/lb/memberv2"
	memberv3 "github.com/enrrou/otc-provider-jet/internal/controller/lb/memberv3"
	monitorv2 "github.com/enrrou/otc-provider-jet/internal/controller/lb/monitorv2"
	monitorv3 "github.com/enrrou/otc-provider-jet/internal/controller/lb/monitorv3"
	policyv3lb "github.com/enrrou/otc-provider-jet/internal/controller/lb/policyv3"
	poolv2 "github.com/enrrou/otc-provider-jet/internal/controller/lb/poolv2"
	poolv3 "github.com/enrrou/otc-provider-jet/internal/controller/lb/poolv3"
	rulev3 "github.com/enrrou/otc-provider-jet/internal/controller/lb/rulev3"
	whitelistv2 "github.com/enrrou/otc-provider-jet/internal/controller/lb/whitelistv2"
	groupv2 "github.com/enrrou/otc-provider-jet/internal/controller/logtank/groupv2"
	topicv2 "github.com/enrrou/otc-provider-jet/internal/controller/logtank/topicv2"
	clusterv1mrs "github.com/enrrou/otc-provider-jet/internal/controller/mrs/clusterv1"
	jobv1 "github.com/enrrou/otc-provider-jet/internal/controller/mrs/jobv1"
	dnatrulev2 "github.com/enrrou/otc-provider-jet/internal/controller/nat/dnatrulev2"
	gatewayv2 "github.com/enrrou/otc-provider-jet/internal/controller/nat/gatewayv2"
	snatrulev2 "github.com/enrrou/otc-provider-jet/internal/controller/nat/snatrulev2"
	floatingipassociatev2networking "github.com/enrrou/otc-provider-jet/internal/controller/networking/floatingipassociatev2"
	floatingipv2networking "github.com/enrrou/otc-provider-jet/internal/controller/networking/floatingipv2"
	networkv2 "github.com/enrrou/otc-provider-jet/internal/controller/networking/networkv2"
	portv2 "github.com/enrrou/otc-provider-jet/internal/controller/networking/portv2"
	routerinterfacev2 "github.com/enrrou/otc-provider-jet/internal/controller/networking/routerinterfacev2"
	routerroutev2 "github.com/enrrou/otc-provider-jet/internal/controller/networking/routerroutev2"
	routerv2 "github.com/enrrou/otc-provider-jet/internal/controller/networking/routerv2"
	secgrouprulev2 "github.com/enrrou/otc-provider-jet/internal/controller/networking/secgrouprulev2"
	secgroupv2networking "github.com/enrrou/otc-provider-jet/internal/controller/networking/secgroupv2"
	subnetv2 "github.com/enrrou/otc-provider-jet/internal/controller/networking/subnetv2"
	vipassociatev2 "github.com/enrrou/otc-provider-jet/internal/controller/networking/vipassociatev2"
	vipv2 "github.com/enrrou/otc-provider-jet/internal/controller/networking/vipv2"
	bucket "github.com/enrrou/otc-provider-jet/internal/controller/obs/bucket"
	bucketobject "github.com/enrrou/otc-provider-jet/internal/controller/obs/bucketobject"
	bucketpolicy "github.com/enrrou/otc-provider-jet/internal/controller/obs/bucketpolicy"
	providerconfig "github.com/enrrou/otc-provider-jet/internal/controller/providerconfig"
	instancev1rds "github.com/enrrou/otc-provider-jet/internal/controller/rds/instancev1"
	instancev3rds "github.com/enrrou/otc-provider-jet/internal/controller/rds/instancev3"
	parametergroupv3 "github.com/enrrou/otc-provider-jet/internal/controller/rds/parametergroupv3"
	readreplicav3 "github.com/enrrou/otc-provider-jet/internal/controller/rds/readreplicav3"
	softwareconfigv1 "github.com/enrrou/otc-provider-jet/internal/controller/rts/softwareconfigv1"
	softwaredeploymentv1 "github.com/enrrou/otc-provider-jet/internal/controller/rts/softwaredeploymentv1"
	stackv1 "github.com/enrrou/otc-provider-jet/internal/controller/rts/stackv1"
	buckets3 "github.com/enrrou/otc-provider-jet/internal/controller/s3/bucket"
	bucketobjects3 "github.com/enrrou/otc-provider-jet/internal/controller/s3/bucketobject"
	bucketpolicys3 "github.com/enrrou/otc-provider-jet/internal/controller/s3/bucketpolicy"
	protectedinstancev1 "github.com/enrrou/otc-provider-jet/internal/controller/sdrs/protectedinstancev1"
	protectiongroupv1 "github.com/enrrou/otc-provider-jet/internal/controller/sdrs/protectiongroupv1"
	filesystemv2 "github.com/enrrou/otc-provider-jet/internal/controller/sfs/filesystemv2"
	shareaccessrulesv2 "github.com/enrrou/otc-provider-jet/internal/controller/sfs/shareaccessrulesv2"
	turbosharev1 "github.com/enrrou/otc-provider-jet/internal/controller/sfs/turbosharev1"
	subscriptionv2 "github.com/enrrou/otc-provider-jet/internal/controller/smn/subscriptionv2"
	topicattributev2 "github.com/enrrou/otc-provider-jet/internal/controller/smn/topicattributev2"
	topicv2smn "github.com/enrrou/otc-provider-jet/internal/controller/smn/topicv2"
	domainv2 "github.com/enrrou/otc-provider-jet/internal/controller/swr/domainv2"
	organizationpermissionsv2 "github.com/enrrou/otc-provider-jet/internal/controller/swr/organizationpermissionsv2"
	organizationv2 "github.com/enrrou/otc-provider-jet/internal/controller/swr/organizationv2"
	repositoryv2 "github.com/enrrou/otc-provider-jet/internal/controller/swr/repositoryv2"
	backuppolicyv2 "github.com/enrrou/otc-provider-jet/internal/controller/vbs/backuppolicyv2"
	backupsharev2 "github.com/enrrou/otc-provider-jet/internal/controller/vbs/backupsharev2"
	backupv2 "github.com/enrrou/otc-provider-jet/internal/controller/vbs/backupv2"
	bandwidthassociatev2 "github.com/enrrou/otc-provider-jet/internal/controller/vpc/bandwidthassociatev2"
	bandwidthv2 "github.com/enrrou/otc-provider-jet/internal/controller/vpc/bandwidthv2"
	eipv1 "github.com/enrrou/otc-provider-jet/internal/controller/vpc/eipv1"
	flowlogv1 "github.com/enrrou/otc-provider-jet/internal/controller/vpc/flowlogv1"
	peeringconnectionaccepterv2 "github.com/enrrou/otc-provider-jet/internal/controller/vpc/peeringconnectionaccepterv2"
	peeringconnectionv2 "github.com/enrrou/otc-provider-jet/internal/controller/vpc/peeringconnectionv2"
	routev2 "github.com/enrrou/otc-provider-jet/internal/controller/vpc/routev2"
	subnetv1 "github.com/enrrou/otc-provider-jet/internal/controller/vpc/subnetv1"
	v1vpc "github.com/enrrou/otc-provider-jet/internal/controller/vpc/v1"
	endpointv1 "github.com/enrrou/otc-provider-jet/internal/controller/vpcep/endpointv1"
	servicev1 "github.com/enrrou/otc-provider-jet/internal/controller/vpcep/servicev1"
	endpointgroupv2 "github.com/enrrou/otc-provider-jet/internal/controller/vpnaas/endpointgroupv2"
	ikepolicyv2 "github.com/enrrou/otc-provider-jet/internal/controller/vpnaas/ikepolicyv2"
	ipsecpolicyv2 "github.com/enrrou/otc-provider-jet/internal/controller/vpnaas/ipsecpolicyv2"
	servicev2 "github.com/enrrou/otc-provider-jet/internal/controller/vpnaas/servicev2"
	siteconnectionv2 "github.com/enrrou/otc-provider-jet/internal/controller/vpnaas/siteconnectionv2"
	alarmnotificationv1 "github.com/enrrou/otc-provider-jet/internal/controller/waf/alarmnotificationv1"
	ccattackprotectionrulev1 "github.com/enrrou/otc-provider-jet/internal/controller/waf/ccattackprotectionrulev1"
	certificatev1 "github.com/enrrou/otc-provider-jet/internal/controller/waf/certificatev1"
	datamaskingrulev1 "github.com/enrrou/otc-provider-jet/internal/controller/waf/datamaskingrulev1"
	domainv1 "github.com/enrrou/otc-provider-jet/internal/controller/waf/domainv1"
	falsealarmmaskingrulev1 "github.com/enrrou/otc-provider-jet/internal/controller/waf/falsealarmmaskingrulev1"
	policyv1waf "github.com/enrrou/otc-provider-jet/internal/controller/waf/policyv1"
	preciseprotectionrulev1 "github.com/enrrou/otc-provider-jet/internal/controller/waf/preciseprotectionrulev1"
	webtamperprotectionrulev1 "github.com/enrrou/otc-provider-jet/internal/controller/waf/webtamperprotectionrulev1"
	whiteblackiprulev1 "github.com/enrrou/otc-provider-jet/internal/controller/waf/whiteblackiprulev1"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		v1.Setup,
		configurationv1.Setup,
		groupv1.Setup,
		policyv1.Setup,
		policyv2.Setup,
		volumev2.Setup,
		policyv3.Setup,
		vaultv3.Setup,
		addonv3.Setup,
		clusterv3.Setup,
		nodepoolv3.Setup,
		nodev3.Setup,
		alarmrule.Setup,
		bmsserverv2.Setup,
		bmstagsv2.Setup,
		floatingipassociatev2.Setup,
		floatingipv2.Setup,
		instancev2.Setup,
		keypairv2.Setup,
		secgroupv2.Setup,
		servergroupv2.Setup,
		volumeattachv2.Setup,
		backuppolicyv1.Setup,
		backupv1.Setup,
		clusterv1.Setup,
		snapshotconfigurationv1.Setup,
		trackerv1.Setup,
		instancev1.Setup,
		instancev3.Setup,
		hostv1.Setup,
		groupv1dms.Setup,
		instancev1dms.Setup,
		queuev1.Setup,
		ptrrecordv2.Setup,
		recordsetv2.Setup,
		zonev2.Setup,
		instancev1ecs.Setup,
		volumev3.Setup,
		firewallgroupv2.Setup,
		policyv2fw.Setup,
		rulev2.Setup,
		agencyv3.Setup,
		credentialv3.Setup,
		groupmembershipv3.Setup,
		groupv3.Setup,
		mappingv3.Setup,
		projectv3.Setup,
		protocolv3.Setup,
		providerv3.Setup,
		roleassignmentv3.Setup,
		rolev3.Setup,
		usergroupmembershipv3.Setup,
		userv3.Setup,
		imageaccessacceptv2.Setup,
		imageaccessv2.Setup,
		imagev2.Setup,
		dataimagev2.Setup,
		imagev2ims.Setup,
		grantv1.Setup,
		keyv1.Setup,
		certificatev2.Setup,
		certificatev3.Setup,
		l7policyv2.Setup,
		l7rulev2.Setup,
		listenerv2.Setup,
		listenerv3.Setup,
		loadbalancerv2.Setup,
		loadbalancerv3.Setup,
		memberv2.Setup,
		memberv3.Setup,
		monitorv2.Setup,
		monitorv3.Setup,
		policyv3lb.Setup,
		poolv2.Setup,
		poolv3.Setup,
		rulev3.Setup,
		whitelistv2.Setup,
		groupv2.Setup,
		topicv2.Setup,
		clusterv1mrs.Setup,
		jobv1.Setup,
		dnatrulev2.Setup,
		gatewayv2.Setup,
		snatrulev2.Setup,
		floatingipassociatev2networking.Setup,
		floatingipv2networking.Setup,
		networkv2.Setup,
		portv2.Setup,
		routerinterfacev2.Setup,
		routerroutev2.Setup,
		routerv2.Setup,
		secgrouprulev2.Setup,
		secgroupv2networking.Setup,
		subnetv2.Setup,
		vipassociatev2.Setup,
		vipv2.Setup,
		bucket.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		providerconfig.Setup,
		instancev1rds.Setup,
		instancev3rds.Setup,
		parametergroupv3.Setup,
		readreplicav3.Setup,
		softwareconfigv1.Setup,
		softwaredeploymentv1.Setup,
		stackv1.Setup,
		buckets3.Setup,
		bucketobjects3.Setup,
		bucketpolicys3.Setup,
		protectedinstancev1.Setup,
		protectiongroupv1.Setup,
		filesystemv2.Setup,
		shareaccessrulesv2.Setup,
		turbosharev1.Setup,
		subscriptionv2.Setup,
		topicattributev2.Setup,
		topicv2smn.Setup,
		domainv2.Setup,
		organizationpermissionsv2.Setup,
		organizationv2.Setup,
		repositoryv2.Setup,
		backuppolicyv2.Setup,
		backupsharev2.Setup,
		backupv2.Setup,
		bandwidthassociatev2.Setup,
		bandwidthv2.Setup,
		eipv1.Setup,
		flowlogv1.Setup,
		peeringconnectionaccepterv2.Setup,
		peeringconnectionv2.Setup,
		routev2.Setup,
		subnetv1.Setup,
		v1vpc.Setup,
		endpointv1.Setup,
		servicev1.Setup,
		endpointgroupv2.Setup,
		ikepolicyv2.Setup,
		ipsecpolicyv2.Setup,
		servicev2.Setup,
		siteconnectionv2.Setup,
		alarmnotificationv1.Setup,
		ccattackprotectionrulev1.Setup,
		certificatev1.Setup,
		datamaskingrulev1.Setup,
		domainv1.Setup,
		falsealarmmaskingrulev1.Setup,
		policyv1waf.Setup,
		preciseprotectionrulev1.Setup,
		webtamperprotectionrulev1.Setup,
		whiteblackiprulev1.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
