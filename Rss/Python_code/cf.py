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
#
# class MovieItem():
#     def __init__(self):
#         self.__movieId = 0
#         self.__rating = 0
#
#     def get_movieId(self):
#         return self.__movieId
#
#     def get_rating(self):
#         return self.__rating
#
#     def set_movieId(self, m):
#         self.__movieId = m
#
#     def set_rating(self, n):
#         self.__rating = n

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


if __name__ == '__main__':
    logging.basicConfig()
    serve()


