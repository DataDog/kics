from pydantic import BaseModel


class Output(BaseModel):
    queryName: str
    descriptionText: str

    def toDict(self):
        return {
            "queryName": self.queryName,
            "descriptionText": self.descriptionText,
        }
