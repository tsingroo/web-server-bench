const std = @import("std");
const zap = @import("zap");

fn on_request(r: zap.Request) anyerror!void {
    if (r.path) |the_path| {
        if (std.mem.eql(u8, the_path, "/ping")) {
            r.sendBody("pong") catch return;
        } else if (std.mem.startsWith(u8, the_path, "/hello/")) {
            // Extract name from /hello/:name
            const name_start = "/hello/".len;
            const name = the_path[name_start..];
            var response_buffer: [256]u8 = undefined;
            const response = std.fmt.bufPrint(&response_buffer, "my name is {s}", .{name}) catch return;
            r.sendBody(response) catch return;
        } else {
            r.sendBody("Not found") catch return;
        }
    } else {
        r.sendBody("No path provided") catch return;
    }
}

pub fn main() !void {
    var listener = zap.HttpListener.init(.{
        .port = 8080,
        .on_request = on_request,
        .log = true,
    });

    std.debug.print("Server listening on http://localhost:8080\n", .{});
    try listener.listen();

    // Start worker threads
    zap.start(.{
        .threads = 2,
        .workers = 2,
    });
}

test "simple test" {
    const gpa = std.testing.allocator;
    var list: std.ArrayList(i32) = .empty;
    defer list.deinit(gpa); // Try commenting this out and see if zig detects the memory leak!
    try list.append(gpa, 42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}

test "fuzz example" {
    const Context = struct {
        fn testOne(context: @This(), input: []const u8) anyerror!void {
            _ = context;
            // Try passing `--fuzz` to `zig build test` and see if it manages to fail this test case!
            try std.testing.expect(!std.mem.eql(u8, "canyoufindme", input));
        }
    };
    try std.testing.fuzz(Context{}, Context.testOne, .{});
}
