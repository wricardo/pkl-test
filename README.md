# Example 06 - pkltest

This is a mini project to showcase how to use `pkl` files to configure your application.

## Idea

`pkl` is a strongly typed configuration framework that can be used to configure your application. In this POC, we have configurations for a multi-service application that has both local and dev environments. Configurations for the local environment are stored in `config/environment/local`, and for the dev environment, in `config/environment/dev`.

**Note:** In a real-world scenario, you would have configurations for dev, staging, and production environments in separate repositories/storage mechanisms.

## How to Run

### Run the Config Service

```bash
go run ./cmd/pkl-test configservice run
```

### Run the Example Service

In another terminal, run the example service (this will exit and print out the configuration):

```bash
go run ./cmd/pkl-test example run
```

### View Configuration in JSON

You can use the following curl command to see the configuration in JSON:

```bash
curl "http://localhost:8080/formation/dev/edu"
```

### View Configuration in `pkl`

You can use the following curl command to see the configuration in `pkl`:

```bash
curl "http://localhost:8080/static/environment/dev/edu/main.pkl"
```

## Nomenclature

### Environment

- local
- dev
- prod

### Service

- zuri
- zasper
- kosmos
- kosmos-ui

### Role

- API Server (GRPC/REST)
- Worker (Temporal)
- Machine Learning Server
- UI Server
- Proxy

### Formation

- Local
- Edu
- Apptness

## Q & A

**Question:** How does the application know the configuration to use?
**Answer:** We'll inject three values: `ENVIRONMENT`, `SERVICE`, and `FORMATION` into the application. The application will use these values to determine the correct configuration to pull down.

**Question:** How do we switch between different environments (local, dev, prod) in our application? Is there a specific command or configuration file to modify?
**Answer:** We use the three values `ENVIRONMENT`, `SERVICE`, and `FORMATION` to determine the configuration to use. We need to change the environment variables and restart/start the application.

**Question:** Could you provide an example of how to load and use a configuration file from the `pkl` framework in one of our services?
**Answer:** Check out `./example/cmd/main.go` for an example of how to load and use a configuration file from the `pkl` framework in one of our services.

**Question:** What is the structure of the configuration files stored in `config/environment/local` and `config/environment/dev`?
**Answer:** The configuration files are stored in a hierarchical structure with the following format:
```bash
config/environment/{ENVIRONMENT}/{FORMATION}/main.pkl
```

**Question:** How do we handle secret management within these configuration files? For instance, where do we store sensitive information like API keys or database passwords?
**Answer:** We can store the sensitive information in a separate file or a secret management tool like Vault. We can then load the sensitive information into the configuration files when the config service starts up or receives a signal to reload.

**Question:** What is the process for updating the configurations for a particular environment, and how do we ensure these updates are propagated to all relevant services?
**Answer:** We can update the configurations for a particular environment by modifying the configuration files stored in the `config/environment/{ENVIRONMENT}/{FORMATION}` directory. We can then use a configuration management tool to deploy the updated configurations.

**Question:** How are roles like 'API Server', 'Worker', 'Machine Learning Server', 'UI Server', and 'Proxy' assigned and configured within this framework?
**Answer:** Roles are assigned for running applications, and they can read the configuration for their service and configure themselves accordingly.

**Question:** If we want to add a new service or a new environment, what steps do we need to follow?
**Answer:** To add a new service or environment, we need to create a new configuration folder in the `config/services/{SERVICE}` directory, then add it to the `config/config.pkl` top-level file.

**Question:** How do we validate that the configurations are correct and the services are properly configured before deploying to an environment?
**Answer:** We can validate the configurations by writing unit tests that check the correctness of the configuration files. Being strongly typed, the `pkl` framework will catch any errors in the configuration files at compile time.

**Question:** How do I see the configuration for a particular service in a particular environment?
**Answer:** Use the curl command:
```bash
curl "http://localhost:8080/formation/dev/edu"
```
