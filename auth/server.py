from concurrent import futures
import logging

import grpc

from proto import auth_pb2
from proto import auth_pb2_grpc


class Auth(auth_pb2_grpc.AuthServicer):

    def Login(self, request, context):
        message="login request was received, for user with email={} and password={}".format(request.email, request.password)
        print(message)
        return auth_pb2.AuthResponse(message=message)

    def Register(self, request, context):
        message = "Register request was received, for user with email={}".format(request.email)
        print(message)
        return auth_pb2.AuthResponse(message=message)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    auth_pb2_grpc.add_AuthServicer_to_server(Auth(), server)
    server.add_insecure_port('[::]:8081')

    print("starting auth service")

    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
