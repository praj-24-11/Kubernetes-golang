from flask import Flask, request, jsonify
import os

app = Flask(__name__)

@app.route('/validate', methods=['POST'])
def validate_chain():
    data = request.get_json()
    mcp_value = data.get('mcp')
    change_request = data.get('change_request')
    # Validate MCP and change management logic
    if mcp_value and change_request and isinstance(mcp_value, str) and isinstance(change_request, str):
        if len(mcp_value) > 0 and len(change_request) > 0:
            return jsonify({"status": "success", "message": "MCP and change validated"}), 200
    return jsonify({"status": "error", "message": "Invalid MCP or change request"}), 400

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=int(os.getenv('PORT', 5000)))