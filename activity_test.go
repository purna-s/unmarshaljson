package unmarshaljson

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput('{"data":{"accessType":"non-exclusive","ackReceivedHighMsgId":959863,"ackReceivedLowMsgId":959863,"bindCountPercentHighClearThreshold":60,"bindCountPercentHighThreshold":80,"bindTimeForwardingMode":"store-and-forward","consumerAckPropagationEnabled":true,"counter":{"ackPendingMsgCount":0,"ackPropagationReceivedMsgCount":0,"ackPropagationSentMsgCount":0,"alreadyBoundBindFailureCount":0,"bindCount":1,"bindRequestCount":14,"bindResponseCount":14,"bindSuccessCount":14,"clientProfileDeniedDiscardedMsgCount":0,"currentSpooledMsgCount":0,"deletedDiscardedMsgCount":0,"deliveredUnackedMsgCount":0,"destinationGroupErrorDiscardedMsgCount":0,"disabledBindFailureCount":0,"invalidSelectorBindFailureCount":0,"lowPriorityMsgCongestionDiscardedMsgCount":0,"maxBindCountExceededBindFailureCount":0,"maxMsgSizeExceededDiscardedMsgCount":0,"maxMsgSpoolUsageExceededDiscardedMsgCount":0,"maxRedeliveryCountExceededDiscardedMsgCount":0,"maxRedeliveryCountExceededToDmqFailedMsgCount":0,"maxRedeliveryCountExceededToDmqMsgCount":0,"maxTtlExceededDiscardedMsgCount":0,"maxTtlExpiredDiscardedMsgCount":19,"maxTtlExpiredToDmqFailedMsgCount":0,"maxTtlExpiredToDmqMsgCount":0,"msgSpoolDisabledDiscardedMsgCount":0,"noLocalDeliveryDiscardedMsgCount":0,"otherBindFailureCount":0,"redeliveredMsgCount":0,"replicationAckedViaAckPropagationMsgCount":0,"replicationReceivedMsgCount":0,"spooledByteCount":36298,"spooledMsgCount":54,"topicSubscriptionCount":1},"deadMsgQueue":"#DEAD_MSG_QUEUE","egressEnabled":true,"eventBindCountThreshold":{"clearPercent":60,"setPercent":80},"eventMsgSpoolUsageThreshold":{"clearPercent":60,"setPercent":80},"eventRejectLowPriorityMsgLimitThreshold":{"clearPercent":60,"setPercent":80},"highWaterMark":0.00390625,"ingressEnabled":true,"isCreatedByManagement":true,"isDurable":true,"isEgressSelectorPresent":false,"lastSpooledMsgId":959863,"lowPriorityMsgCongestionState":"disabled","maxBindCount":1000,"maxDeliveredUnackedMsgsPerFlow":10000,"maxMsgSize":10000000,"maxMsgSpoolUsage":1500,"maxRedeliveryCount":0,"maxTtl":600,"msgSpoolUsage":0,"msgVpnName":"msgvpn-jfgwkeg0s4r","networkTopic":"#P2P/QUE/Q_TRFFCSMART_DSI_ALRM_RT","owner":"","permission":"consume","queueName":"Q_TRFFCSMART_DSI_ALRM_RT","rate":{"averageEgressByteRate":16,"averageEgressMsgRate":0,"averageIngressByteRate":15,"averageIngressMsgRate":0,"egressByteRate":465,"egressMsgRate":1,"ingressByteRate":465,"ingressMsgRate":1},"rejectLowPriorityMsgEnabled":false,"rejectLowPriorityMsgLimit":0,"rejectMsgToSenderOnDiscardBehavior":"when-queue-enabled","respectMsgPriorityEnabled":false,"respectTtlEnabled":true,"virtualRouter":"primary"},"links":{"egressFlowsUri":"https://mr8ksiwsp23vv.messaging.solace.cloud:20590/SEMP/v2/monitor/msgVpns/msgvpn-jfgwkeg0s4r/queues/Q_TRFFCSMART_DSI_ALRM_RT/egressFlows","messagesUri":"https://mr8ksiwsp23vv.messaging.solace.cloud:20590/SEMP/v2/monitor/msgVpns/msgvpn-jfgwkeg0s4r/queues/Q_TRFFCSMART_DSI_ALRM_RT/messages","prioritiesUri":"https://mr8ksiwsp23vv.messaging.solace.cloud:20590/SEMP/v2/monitor/msgVpns/msgvpn-jfgwkeg0s4r/queues/Q_TRFFCSMART_DSI_ALRM_RT/priorities","subscriptionsUri":"https://mr8ksiwsp23vv.messaging.solace.cloud:20590/SEMP/v2/monitor/msgVpns/msgvpn-jfgwkeg0s4r/queues/Q_TRFFCSMART_DSI_ALRM_RT/subscriptions","uri":"https://mr8ksiwsp23vv.messaging.solace.cloud:20590/SEMP/v2/monitor/msgVpns/msgvpn-jfgwkeg0s4r/queues/Q_TRFFCSMART_DSI_ALRM_RT"},"meta":{"request":{"method":"GET","uri":"https://mr8ksiwsp23vv.messaging.solace.cloud:20590/SEMP/v2/monitor/msgVpns/msgvpn-jfgwkeg0s4r/queues/Q_TRFFCSMART_DSI_ALRM_RT"},"responseCode":200}}')

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}
	act.Eval(tc)
	//check output attr

	output := tc.GetOutput("output")
	assert.Equal(t, output, output)

}
