# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import recommed_rpc_pb2 as recommed__rpc__pb2


class RecommendStub(object):
  """python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. recommed_rpc.proto
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Recommend = channel.unary_unary(
        '/RecommendPro.Recommend/Recommend',
        request_serializer=recommed__rpc__pb2.RecommendReq.SerializeToString,
        response_deserializer=recommed__rpc__pb2.RecommendReply.FromString,
        )


class RecommendServicer(object):
  """python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. recommed_rpc.proto
  """

  def Recommend(self, request, context):
    """Sends a greeting
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RecommendServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Recommend': grpc.unary_unary_rpc_method_handler(
          servicer.Recommend,
          request_deserializer=recommed__rpc__pb2.RecommendReq.FromString,
          response_serializer=recommed__rpc__pb2.RecommendReply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'RecommendPro.Recommend', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
