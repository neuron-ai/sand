from subprocess import run
from os import get_terminal_size

print()
print("=" * get_terminal_size().columns)
print("Please ensure that you have read the CODE_OF_CONDUCT.md file and the CONTRIBUTING.md contribution guidelines.".center(get_terminal_size().columns, " "))
print("=" * get_terminal_size().columns)
print("\nAssigning git hook path...")
run("git config core.hooksPath .githooks")

print()
print(" All Done! ".center(get_terminal_size().columns, "="))