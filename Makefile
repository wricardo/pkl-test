clean: 
	find . -name '*.pkl.go' -delete

generate: clean
	# pkl-gen-go ./environment/dev-edu/main.pkl
	pkl-gen-go formation.pkl
	# hack
	# mv bitbucket.org/zetaactions/pkl-test/* .
	# rm -rf bitbucket.org
eval:
	pkl eval --allowed-resources "http://*,prop:*" --format json environment/dev/edu/main.pkl
