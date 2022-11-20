from tasket.models import RestRequest, RestResponse


class TaskResource:
    def __init__(self):
        pass

    def get_task(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)

    def get_tasks(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)

    def get_task_logs(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)

    def stop_task(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)

    def update_task(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)

    def create_task(self, req: RestRequest) -> RestResponse:
        return RestResponse("", 200)
