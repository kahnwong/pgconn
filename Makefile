setup:
	pyenv virtualenv 3.11.4 pgconn
	pyenv local pgconn
	pip install -e .
