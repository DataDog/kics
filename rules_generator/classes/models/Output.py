from pydantic import BaseModel


class QueryAndDescriptionOutput(BaseModel):
    queryName: str
    descriptionText: str

    def toDict(self):
        return {
            "queryName": self.queryName,
            "descriptionText": self.descriptionText,
        }


class DescriptionOutput(BaseModel):
    descriptionText: str

    def toDict(self):
        return {
            "descriptionText": self.descriptionText,
        }
