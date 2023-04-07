
from flask_marshmallow import Marshmallow
from ..models.produto import Produtos

ma = Marshmallow()

class ProdutoSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Produtos
        load_instance = True
