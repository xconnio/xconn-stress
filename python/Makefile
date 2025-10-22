install_uv:
	@if ! command -v uv >/dev/null 2>&1; then \
  		curl -LsSf https://astral.sh/uv/install.sh | sh; \
  	fi

setup:
	make install_uv
	uv venv
	uv pip install .[test] -U

format:
	./.venv/bin/ruff format .

lint:
	./.venv/bin/ruff check .

imager:
	./.venv/bin/python3 image-publisher.py

clean:
	rm -rf build *.egg-info
