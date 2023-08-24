setup:
	pyenv virtualenv 3.11.4 pg-conn-cli
	pyenv local pg-conn-cli
	pip install -e .
