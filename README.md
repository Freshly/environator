# Environator
Takes a few YAML file inputs and generates some YAML files as outputs.

## What is it?
Environator was built to deal with the complexity of creating a common base of environment variables for the Freshly monolith.

The goal of the tool is to have a base value set based off production, and create two override files with staging values. In an ideal world, most of the keys are the same between all environments.

For example:
`prod.yaml` contains all of the keys/values for the apps environment. We want to generate a staging config based on this set of values.