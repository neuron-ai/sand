from subprocess import run, PIPE

git = run(["git", "status"], stdout=PIPE, stderr=PIPE)
output = git.stdout.decode("utf-8")

start = output.find("Changes to be committed:") + len("Changes to be committed:")
end = output.find('Changes not staged for commit', start)

if end == -1:
	end = output.find("Untracked files", start)

output = output[start:end]

for line in output.splitlines():
	if line.endswith(".go"):
		filepath = line.split("  ")[-1]
		run(f"go fmt {filepath}")
