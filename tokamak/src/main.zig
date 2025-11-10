const std = @import("std");
const tk = @import("tokamak");

const routes: []const tk.Route = &.{
    .get("/ping", ping),
    .get("/hello/:name", hello),
};

fn ping() []const u8 {
    return "pong";
}

fn hello(arena: std.mem.Allocator, name: []const u8) ![]const u8 {
    return std.fmt.allocPrint(arena, "my name is {s}", .{name});
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    var server = try tk.Server.init(allocator, routes, .{ .listen = .{ .port = 9999 } });
    defer server.deinit();
    try server.start();
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
