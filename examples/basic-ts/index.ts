import * as pulumi from "@pulumi/pulumi";
import * as komodor from "@pulumi/komodor";

const policy = new komodor.Policy("my-policy", { name: "my-policy" });

export const policyId = policy.id;
