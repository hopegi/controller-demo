package scheme

import (
	ecsv1 "controller-demo/api/v1"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(scheme)
var ParameterCodec = runtime.NewParameterCodec(scheme)

//var localSchemeBuilder = runtime.SchemeBuilder{
//	ecsv1.AddToScheme,
//}
//
//var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	metav1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	if err := AddToScheme(scheme); err != nil {
		fmt.Println("error to AddToScheme ", err)
	}
}

func AddToScheme(scheme *runtime.Scheme) error {
	return ecsv1.AddToScheme(scheme)
}
