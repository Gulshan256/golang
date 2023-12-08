from flask import Flask, request, jsonify

app = Flask(__name__)

# Sample data (you can replace this with a database)
data = {
    1: {'name': 'John', 'age': 25},
    2: {'name': 'Alice', 'age': 30},
    3: {'name': 'Bob', 'age': 22}
}

# GET request to retrieve all data
@app.route('/api/data', methods=['GET'])
def get_all_data():
    return jsonify(data)

# GET request to retrieve a specific item by ID
@app.route('/api/data/<int:item_id>', methods=['GET'])
def get_data(item_id):
    if item_id in data:
        return jsonify(data[item_id])
    else:
        return jsonify({'error': 'Item not found'}), 404

# POST request to add new data
@app.route('/api/data', methods=['POST'])
def add_data():
    #  check if data
    
    new_item = request.get_json()
    
    
    item_id = max(data.keys()) + 1
    data[item_id] = new_item
    return jsonify({'id': item_id}), 201


# POST request to add new data with form data
@app.route('/api/form-data', methods=['POST'])
def add_form_data():
    try:
        new_item = {
            'name': request.form['name'],
            'age': int(request.form['age'])
        }
        item_id = max(data.keys()) + 1
        data[item_id] = new_item
        return jsonify({'id': item_id,'data': data}), 201
    except KeyError:
        return jsonify({'key error': 'Invalid form data'}), 400
    except ValueError:
        return jsonify({'vakue error': 'Invalid age value'}), 400


# PUT request to update existing data
@app.route('/api/data/<int:item_id>', methods=['PUT'])
def update_data(item_id):
    if item_id in data:
        updated_item = request.get_json()
        data[item_id] = updated_item
        return jsonify({'message': 'Item updated successfully'})
    else:
        return jsonify({'error': 'Item not found'}), 404

# DELETE request to remove an item by ID
@app.route('/api/data/<int:item_id>', methods=['DELETE'])
def delete_data(item_id):
    if item_id in data:
        del data[item_id]
        return jsonify({'message': 'Item deleted successfully'})
    else:
        return jsonify({'error': 'Item not found'}), 404

if __name__ == '__main__':
    app.run(debug=True)

    """
    To get all data: GET http://127.0.0.1:5000/api/data
    To get a specific item: GET http://127.0.0.1:5000/api/data/1
    To add new data: POST http://127.0.0.1:5000/api/data with a JSON payload
    To add new data: POST http://127.0.0.1:5000/api/form-data with form data
    To update existing data: PUT http://127.0.0.1:5000/api/data/1 with a JSON payload
    delete an item: DELETE http://127.0.0.1:5000/api/data/1

    """
    
