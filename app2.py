from flask import Flask, request, jsonify

app = Flask(__name__)

# In-memory storage for Beverages
beverages = {}
next_id = 1


@app.route('/beverages', methods=['POST'])
def create_beverage():
    """Create a new beverage."""
    global next_id
    data = request.get_json()

    # Validate input
    if not data or 'name' not in data or 'type' not in data:
        return jsonify({"error": "Invalid input. 'name' and 'type' are required."}), 400

    # Add to storage
    beverage = {"id": next_id, "name": data['name'], "type": data['type']}
    beverages[next_id] = beverage
    next_id += 1

    return jsonify(beverage), 201


@app.route('/beverages', methods=['GET'])
def get_all_beverages():
    """Retrieve all beverages."""
    return jsonify(list(beverages.values())), 200


@app.route('/beverages/<int:beverage_id>', methods=['GET'])
def get_beverage(beverage_id):
    """Retrieve a single beverage by ID."""
    beverage = beverages.get(beverage_id)
    if not beverage:
        return jsonify({"error": f"Beverage with ID {beverage_id} not found."}), 404
    return jsonify(beverage), 200


@app.route('/beverages/<int:beverage_id>', methods=['PUT'])
def update_beverage(beverage_id):
    """Update an existing beverage."""
    data = request.get_json()

    # Validate input
    if not data or 'name' not in data or 'type' not in data:
        return jsonify({"error": "Invalid input. 'name' and 'type' are required."}), 400

    # Check if the beverage exists
    if beverage_id not in beverages:
        return jsonify({"error": f"Beverage with ID {beverage_id} not found."}), 404

    # Update the beverage
    beverages[beverage_id]['name'] = data['name']
    beverages[beverage_id]['type'] = data['type']

    return jsonify(beverages[beverage_id]), 200


@app.route('/beverages/<int:beverage_id>', methods=['DELETE'])
def delete_beverage(beverage_id):
    """Delete a beverage."""
    if beverage_id not in beverages:
        return jsonify({"error": f"Beverage with ID {beverage_id} not found."}), 404

    del beverages[beverage_id]
    return jsonify({"message": f"Beverage with ID {beverage_id} deleted."}), 200


if __name__ == '__main__':
    app.run(debug=True)
