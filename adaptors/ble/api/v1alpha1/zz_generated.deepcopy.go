// +build !ignore_autogenerated

/*

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BluetoothDataConverter) DeepCopyInto(out *BluetoothDataConverter) {
	*out = *in
	if in.OrderOfOperations != nil {
		in, out := &in.OrderOfOperations, &out.OrderOfOperations
		*out = make([]BluetoothOperations, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BluetoothDataConverter.
func (in *BluetoothDataConverter) DeepCopy() *BluetoothDataConverter {
	if in == nil {
		return nil
	}
	out := new(BluetoothDataConverter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BluetoothDevice) DeepCopyInto(out *BluetoothDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BluetoothDevice.
func (in *BluetoothDevice) DeepCopy() *BluetoothDevice {
	if in == nil {
		return nil
	}
	out := new(BluetoothDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BluetoothDevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BluetoothDeviceList) DeepCopyInto(out *BluetoothDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BluetoothDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BluetoothDeviceList.
func (in *BluetoothDeviceList) DeepCopy() *BluetoothDeviceList {
	if in == nil {
		return nil
	}
	out := new(BluetoothDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BluetoothDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BluetoothDeviceSpec) DeepCopyInto(out *BluetoothDeviceSpec) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]DeviceProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BluetoothDeviceSpec.
func (in *BluetoothDeviceSpec) DeepCopy() *BluetoothDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(BluetoothDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BluetoothDeviceStatus) DeepCopyInto(out *BluetoothDeviceStatus) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]StatusProperties, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BluetoothDeviceStatus.
func (in *BluetoothDeviceStatus) DeepCopy() *BluetoothDeviceStatus {
	if in == nil {
		return nil
	}
	out := new(BluetoothDeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BluetoothOperations) DeepCopyInto(out *BluetoothOperations) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BluetoothOperations.
func (in *BluetoothOperations) DeepCopy() *BluetoothOperations {
	if in == nil {
		return nil
	}
	out := new(BluetoothOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceProperty) DeepCopyInto(out *DeviceProperty) {
	*out = *in
	in.Visitor.DeepCopyInto(&out.Visitor)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceProperty.
func (in *DeviceProperty) DeepCopy() *DeviceProperty {
	if in == nil {
		return nil
	}
	out := new(DeviceProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyVisitor) DeepCopyInto(out *PropertyVisitor) {
	*out = *in
	if in.DataWriteTo != nil {
		in, out := &in.DataWriteTo, &out.DataWriteTo
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	in.BluetoothDataConverter.DeepCopyInto(&out.BluetoothDataConverter)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyVisitor.
func (in *PropertyVisitor) DeepCopy() *PropertyVisitor {
	if in == nil {
		return nil
	}
	out := new(PropertyVisitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusProperties) DeepCopyInto(out *StatusProperties) {
	*out = *in
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusProperties.
func (in *StatusProperties) DeepCopy() *StatusProperties {
	if in == nil {
		return nil
	}
	out := new(StatusProperties)
	in.DeepCopyInto(out)
	return out
}
