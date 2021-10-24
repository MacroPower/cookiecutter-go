#!/usr/bin/env python
import os
import shutil
import yaml

MANIFEST = "manifest.yaml"


def delete_resources_for_disabled_features():
    with open(MANIFEST) as manifest_file:
        manifest = yaml.safe_load(manifest_file)
        for feature in manifest["features"]:
            print("handling feature {}".format(feature["name"]))
            if feature["enabled"] != "y":
                print("removing disabled feature {}".format(feature["name"]))
                for resource in feature["resources"]:
                    delete_resource(resource)
    print("cleanup complete, removing manifest")
    delete_resource(MANIFEST)


def delete_resource(resource):
    if os.path.isfile(resource):
        print("removing file: {}".format(resource))
        os.remove(resource)
    elif os.path.isdir(resource):
        print("removing directory: {}".format(resource))
        shutil.rmtree(resource)


if __name__ == "__main__":
    delete_resources_for_disabled_features()
