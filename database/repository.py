from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from models.produto import Base, Produtos

DB = "mysql+mysqlconnector://root:root@localhost:3306/adega"

class Repository:
    def __init__(self):
        engine = create_engine(DB)
        Base.metadata.create_all(engine)
        self.Session = sessionmaker(bind=engine)

    def create_produto(self, nome, descricao, valor, quantidade):
        session = self.Session()
        

    def get_produtos(self):
        session = self.Session()
        