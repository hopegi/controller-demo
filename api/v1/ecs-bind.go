package v1

import (
	//"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
	//"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (

	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "hopegi.com", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	//SchemeBuilder = runtime.NewSchemeBuilder(AddKnownTypes)

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	//Scheme = runtime.NewScheme()
	//Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {

	SchemeBuilder.Register(&EcsBinding{}, &EcsBindingList{})
	//fmt.Println("type has regist into scheme")

	//metav1.AddToGroupVersion(Scheme,GroupVersion)
	//AddToScheme(Scheme)

}

//func AddKnownTypes(scheme *runtime.Scheme) error {
//	scheme.AddKnownTypes(GroupVersion,&EcsBinding{},&EcsBindingList{})
//	metav1.AddToGroupVersion(scheme,GroupVersion)
//	return nil
//}

type EcsBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EcsBindSpec   `json:"spec,omitempty"`
	Status EcsBindStatus `json:"status,omitempty"`
}

type EcsBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []EcsBinding `json:"items"`
}

type EcsBindSpec struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	NodeName string `json:"node_name"`
	InnerIp  string `json:"inner_ip"`
}

type EcsBindStatus struct {
}

func (in *EcsBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *EcsBinding) DeepCopy() *EcsBinding {
	if in == nil {
		return nil
	}
	out := new(EcsBinding)
	in.DeepCopyInto(out)
	return out
}

func (in *EcsBinding) DeepCopyInto(out *EcsBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

func (in *EcsBindSpec) DeepCopyInto(out *EcsBindSpec) {
	*out = *in
}

func (in *EcsBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *EcsBindingList) DeepCopy() *EcsBindingList {
	if in == nil {
		return nil
	}
	out := new(EcsBindingList)
	in.DeepCopyInto(out)
	return out
}

func (in *EcsBindingList) DeepCopyInto(out *EcsBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EcsBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}
