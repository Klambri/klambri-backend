#!/bin/bash
python3 -m venv .venv --prompt ansible
.venv/bin/pip --quiet install ansible ansible-lint
echo -e "Using: \033[0;36msource .venv/bin/activate\033[0m"