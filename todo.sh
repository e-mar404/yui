#! /usr/bin/env bash

grep "TODO" -rni --color --exclude="todo.sh" --exclude-dir=".git" ./
