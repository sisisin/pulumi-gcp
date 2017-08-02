// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Cluster extends lumi.NamedResource implements ClusterArgs {
    public readonly additionalZones: string[];
    public readonly addonsConfig?: { horizontalPodAutoscaling?: { disabled?: boolean }[], httpLoadBalancing?: { disabled?: boolean }[] }[];
    public readonly clusterIpv4Cidr: string;
    public readonly description?: string;
    public /*out*/ readonly endpoint: string;
    public readonly initialNodeCount?: number;
    public /*out*/ readonly instanceGroupUrls: string[];
    public readonly loggingService: string;
    public readonly masterAuth: { clientCertificate: string, clientKey: string, clusterCaCertificate: string, password: string, username: string }[];
    public readonly monitoringService: string;
    public readonly clusterName?: string;
    public readonly network?: string;
    public readonly nodeConfig: { diskSizeGb: number, imageType: string, labels?: {[key: string]: string}, localSsdCount: number, machineType: string, metadata?: {[key: string]: string}, oauthScopes: string[], serviceAccount: string, tags?: string[] }[];
    public readonly nodePool: { initialNodeCount: number, name: string, namePrefix?: string }[];
    public readonly nodeVersion: string;
    public readonly project?: string;
    public readonly subnetwork?: string;
    public readonly zone: string;

    constructor(name: string, args: ClusterArgs) {
        super(name);
        this.additionalZones = args.additionalZones;
        this.addonsConfig = args.addonsConfig;
        this.clusterIpv4Cidr = args.clusterIpv4Cidr;
        this.description = args.description;
        this.initialNodeCount = args.initialNodeCount;
        this.loggingService = args.loggingService;
        this.masterAuth = args.masterAuth;
        this.monitoringService = args.monitoringService;
        this.clusterName = args.clusterName;
        this.network = args.network;
        this.nodeConfig = args.nodeConfig;
        this.nodePool = args.nodePool;
        this.nodeVersion = args.nodeVersion;
        this.project = args.project;
        this.subnetwork = args.subnetwork;
        if (lumirt.defaultIfComputed(args.zone, "") === undefined) {
            throw new Error("Property argument 'zone' is required, but was missing");
        }
        this.zone = args.zone;
    }
}

export interface ClusterArgs {
    readonly additionalZones?: string[];
    readonly addonsConfig?: { horizontalPodAutoscaling?: { disabled?: boolean }[], httpLoadBalancing?: { disabled?: boolean }[] }[];
    readonly clusterIpv4Cidr?: string;
    readonly description?: string;
    readonly initialNodeCount?: number;
    readonly loggingService?: string;
    readonly masterAuth?: { clientCertificate: string, clientKey: string, clusterCaCertificate: string, password: string, username: string }[];
    readonly monitoringService?: string;
    readonly clusterName?: string;
    readonly network?: string;
    readonly nodeConfig?: { diskSizeGb: number, imageType: string, labels?: {[key: string]: string}, localSsdCount: number, machineType: string, metadata?: {[key: string]: string}, oauthScopes: string[], serviceAccount: string, tags?: string[] }[];
    readonly nodePool?: { initialNodeCount: number, name: string, namePrefix?: string }[];
    readonly nodeVersion?: string;
    readonly project?: string;
    readonly subnetwork?: string;
    readonly zone: string;
}
