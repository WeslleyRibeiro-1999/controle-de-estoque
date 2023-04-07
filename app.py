from flask import Flask, jsonify, request
from src.utils.ma import ProdutoSchema
from src.database.repository import Repository

app = Flask('__name__')
repository = Repository()

produto_schema = ProdutoSchema()

@app.route('/produto', methods=['POST'])
def create_produto():
    nome = request.json['nome']
    descricao = request.json['descricao']
    valor = request.json['valor']
    quantidade = request.json['quantidade']

    repository.create_produto(nome, descricao, valor, quantidade)
    
    return jsonify('created'), 201


if __name__ == '__main__':
    app.run(debug=True, port=8080)