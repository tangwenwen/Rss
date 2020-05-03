import time
import logging
import grpc
from concurrent import futures
# python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. recommed_rpc.proto
import recommed_rpc_pb2
import recommed_rpc_pb2_grpc
import user_cf
import os

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class Recommend(recommed_rpc_pb2_grpc.RecommendServicer):

    def Recommend(self, request, context):
        ratingfile = os.path.join('./data/ml-1m', 'ratings.dat')
        usercf = user_cf.UserBasedCF()
        usercf.generate_dataset(ratingfile)
        usercf.calc_user_sim()
        items = usercf.recommend(str(request.userId))
        itemList = []
        for item in items:
            tempDic = {}
            tempDic['moiveId'] = int(item[0])
            tempDic['rating'] = float(item[1])
            itemList.append(tempDic)
        return recommed_rpc_pb2.RecommendReply(recommendItem=itemList)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    recommed_rpc_pb2_grpc.add_RecommendServicer_to_server(Recommend(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

def initServer():
    print("endding")

if __name__ == '__main__':
    logging.basicConfig()
    initServer()
    serve()


