clean: 
	find . -name '*.pkl.go' -delete

generate: clean
	pkl-gen-go ./config/config.pkl

eval:
	pkl eval --allowed-resources "http://*,prop:*" --format json config/environment/dev/edu/main.pkl
