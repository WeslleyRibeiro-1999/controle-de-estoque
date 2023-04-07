from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from ..models.produto import Base, Produtos
from .config import DB_URL

# DB = "mysql+mysqlconnector://root:root@localhost:3306/adega"

class Repository:
    def __init__(self):
        engine = create_engine(DB_URL)
        Base.metadata.create_all(engine)
        self.Session = sessionmaker(bind=engine)

    def create_produto(self, nome, descricao, valor, quantidade):
        session = self.Session()
        produto = Produtos(nome=nome, descricao=descricao, valor=valor, quantidade=quantidade)
        session.add(produto)
        session.commit()
        session.close

    def get_produtos(self):
        session = self.Session()


repository = Repository()

repository.create_produto('MACA', 'ALGO NOVO', 10.50, 15)