from tasket.models import RestRequest, RestResponse


class AuthResource:
    def __init__(self):
        pass

    def create_token(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)

    def invalidate_token(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)
