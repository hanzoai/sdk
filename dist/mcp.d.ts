/**
 * MCP (Model Context Protocol) server
 */
export interface MCPTool {
    name: string;
    handler: (action: string, params: any) => Promise<any>;
}
export declare class MCPServer {
    private transport;
    private tools;
    constructor(transport?: string);
    private registerDefaultTools;
    registerTool(name: string, handler: (action: string, params: any) => Promise<any>): void;
    private filesystemTool;
    private shellTool;
    private memoryTool;
    serve(port?: number): Promise<void>;
    listTools(): string[];
}
//# sourceMappingURL=mcp.d.ts.map