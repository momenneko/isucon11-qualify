# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: isuxportal/services/dcim/instance_management.proto for package 'isuxportal.proto.services.dcim'

require 'grpc'
require 'isuxportal/services/dcim/instance_management_pb'

module Isuxportal
  module Proto
    module Services
      module Dcim
        module InstanceManagement
          class Service

            include GRPC::GenericService

            self.marshal_class_method = :encode
            self.unmarshal_class_method = :decode
            self.service_name = 'isuxportal.proto.services.dcim.InstanceManagement'

            rpc :InformInstanceStateUpdate, Isuxportal::Proto::Services::Dcim::InformInstanceStateUpdateRequest, Isuxportal::Proto::Services::Dcim::InformInstanceStateUpdateResponse
          end

          Stub = Service.rpc_stub_class
        end
      end
    end
  end
end