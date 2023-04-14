
from marshmallow_sqlalchemy import SQLAlchemyAutoSchema
from ..models.produto import Produtos


class ProdutoSchema(SQLAlchemyAutoSchema):
    class Meta:
        model = Produtos
        load_instance = True
