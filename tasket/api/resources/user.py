from tasket.models import RestRequest, RestResponse


class UserResource:
    def get_user(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)

    def create_user(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)

    def delete_user(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)
