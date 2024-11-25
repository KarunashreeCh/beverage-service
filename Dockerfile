# Use Python base image
FROM python:3.11-slim

# Set the working directory
WORKDIR /app

# Copy files
COPY app.py /app/
COPY requirements.txt /app/

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Expose the application port
EXPOSE 5000

# Start the service
CMD ["python", "app.py"]

