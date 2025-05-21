# Logging Demo

Demo app showcasing centralized application logging with the Elastic stack (Elastic Common Schema).

The demo app uses zap logging library and ECS plugin.

To run:

1. Edit `filebeat.docker.yml` file, `output.elasticsearch` entry to point it at your Elasticsearch deployment.

2. Run `docker-compose up -d`.

3. Explore logs in your Elasticsearch/Kibana.
