from flask import Flask, request, jsonify

app = Flask(__name__)

# In-memory database for beverages
beverages = {}
next_id = 1

# Create a beverage
@app.route('/beverages', methods=['POST'])
def create_beverage():
    global next_id
    data = request.json
    if not all(k in data for k in ('name', 'quantity', 'cost')):
        return jsonify({'error': 'Invalid data'}), 400

    beverage = {
        'id': next_id,
        'name': data['name'],
        'quantity': data['quantity'],
        'cost': data['cost']
    }
    beverages[next_id] = beverage
    next_id += 1
    return jsonify(beverage), 201

# Read all beverages
@app.route('/beverages', methods=['GET'])
def get_beverages():
    return jsonify(list(beverages.values())), 200

# Read a single beverage
@app.route('/beverages/<int:beverage_id>', methods=['GET'])
def get_beverage(beverage_id):
    beverage = beverages.get(beverage_id)
    if not beverage:
        return jsonify({'error': 'Beverage not found'}), 404
    return jsonify(beverage), 200

# Update a beverage
@app.route('/beverages/<int:beverage_id>', methods=['PUT'])
def update_beverage(beverage_id):
    data = request.json
    beverage = beverages.get(beverage_id)
    if not beverage:
        return jsonify({'error': 'Beverage not found'}), 404

    beverage['name'] = data.get('name', beverage['name'])
    beverage['quantity'] = data.get('quantity', beverage['quantity'])
    beverage['cost'] = data.get('cost', beverage['cost'])
    return jsonify(beverage), 200

# Delete a beverage
@app.route('/beverages/<int:beverage_id>', methods=['DELETE'])
def delete_beverage(beverage_id):
    if beverage_id not in beverages:
        return jsonify({'error': 'Beverage not found'}), 404
    del beverages[beverage_id]
    return '', 204

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)

