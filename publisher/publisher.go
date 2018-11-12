package publisher

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrcutils -lstd_msgs__rosidl_typesupport_c
// #include <rosidl_generator_c/message_type_support_struct.h>
// #include "rcl/rcl.h"
// #include <std_msgs/msg/string.h>
// #include  <std_msgs/msg/string__functions.h>
//#define ZERO_ALLOCATE(s) \
//  rcl_get_default_allocator().zero_allocate(s, 1, rcl_get_default_allocator().state)
//
// int publish (const rcl_publisher_t * publisher, void* msg){
//		int retValue = rcl_publish(publisher,msg);
//		return retValue;
//}
import "C"
import (
	"rclgo/node"
	"rclgo/types"
	"unsafe"
)

type Publisher struct {
	RCLPublisher *C.rcl_publisher_t
}

type PublisherOptions struct {
	RCLPublisherOptions *C.rcl_publisher_options_t
}

func GetZeroInitializedPublisher() Publisher {
	zeroPublisher := C.rcl_get_zero_initialized_publisher()
	return Publisher{&zeroPublisher}
}

func GetPublisherDefaultOptions() PublisherOptions {
	defOpts := C.rcl_publisher_get_default_options()
	return PublisherOptions{&defOpts}
}

func PublisherInit(publisher Publisher, publisherOptions PublisherOptions, node node.Node, topicName string, msg types.MessageTypeSupport) types.RCLRetT {

	tName := C.CString(topicName)
	defer C.free(unsafe.Pointer(tName))

	return types.RCLRetT(C.rcl_publisher_init(publisher.RCLPublisher,
		(*C.struct_rcl_node_t)(unsafe.Pointer(node.RCLNode)),
		(*C.rosidl_message_type_support_t)(unsafe.Pointer(msg.ROSIdlMessageTypeSupport)),
		tName,
		publisherOptions.RCLPublisherOptions))

}

func Publish(publisher Publisher, msg types.Message) types.RCLRetT {
	retvalue := C.publish(publisher.RCLPublisher, msg.GetMessage())
	return types.RCLRetT(retvalue)
}

// func SubscriptionFini(subscription Subscription, node node.Node) types.RCLRetT {

// 	return types.RCLRetT(C.rcl_subscription_fini(subscription.RCLSubscription,
// 		(*C.struct_rcl_node_t)(unsafe.Pointer(node.RCLNode))))

// }
