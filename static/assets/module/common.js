(() => {
    var n = {
        351: () => {
            layui.define(["form", "layer", "laydate", "upload", "element", "base"], (function (n) {
                "use strict";
                var e = layui.form, t = void 0 === parent.layer ? layui.layer : top.layer, a = layui.laydate,
                    i = layui.upload, o = (layui.element, layui.base), r = layui.$, l = {
                        edit: function (n, e = 0, t = 0, a = 0, i = [], r = null, s = !1) {
                            var c = e > 0 ? "修改" : "新增";
                            o.isEmpty(n) ? c += "内容" : c += n;
                            var u = cUrl + "/edit?id=" + e;
                            if (Array.isArray(i)) for (var f in i) u += "&" + i[f];
                            l.showWin(c, u, t, a, i, 2, [], (function (n, e) {
                                r && r(n, e)
                            }), s)
                        }, detail: function (n, e, t = 0, a = 0, i = !1) {
                            var o = cUrl + "/detail?id=" + e;
                            l.showWin(n + "详情", o, t, a, [], 2, [], null, i)
                        }, cache: function (n) {
                            var e = cUrl + "/cache";
                            l.ajaxPost(e, {id: n}, (function (n, e) {
                            }))
                        }, copy: function (n, e, t = 0, a = 0) {
                            var i = cUrl + "/copy?id=" + e;
                            l.showWin(n + "复制", i, t, a)
                        }, delete: function (n, e = null) {
                            t.confirm("您确定要删除吗？删除后将无法恢复！", {
                                icon: 3,
                                skin: "layer-ext-moon",
                                btn: ["确认", "取消"]
                            }, (function (a) {
                                var i = cUrl + "/delete?id=" + n;
                                console.log(i), l.ajaxPost(i, {}, (function (n, i) {
                                    e && (t.close(a), e(n, i))
                                }), "正在删除。。。")
                            }))
                        }, batchFunc: function (n, e = null) {
                            var a = n.url, i = n.title, o = (n.form, n.confirm || !1), r = n.show_tips || "处理中...",
                                s = n.data || [], c = n.param || [], u = n.type || "POST";
                            if ("导出数据" != i && 0 == s.length) return t.msg("请选择数据", {icon: 5}), !1;
                            var f = [];
                            for (var d in s) f.push(s[d].Id);
                            var m = f.join(","), p = {};
                            if (p.id = m, Array.isArray(c)) for (var d in c) {
                                var y = c[d].split("=");
                                p[y[0]] = y[1]
                            }
                            console.log(p), o ? t.confirm("您确定要【" + i + "】选中的数据吗？", {
                                icon: 3,
                                title: "提示信息"
                            }, (function (n) {
                                "POST" == u ? a.indexOf("/delete") >= 0 ? l.ajaxPost(a + "?id=" + m, {}, e, r) : l.ajaxPost(a, p, e, r) : l.ajaxGet(a + "/" + m, {}, e, r)
                            })) : "POST" == u ? l.ajaxPost(a, p, e, r) : l.ajaxGet(a + "/" + m, {}, e, r)
                        }, verify: function () {
                            e.verify({
                                number: [/^[0-9]*$/, "请输入数字"], username: function (n, e) {
                                    return new RegExp("^[a-zA-Z0-9_一-龥\\s·]+$").test(n) ? /(^\_)|(\__)|(\_+$)/.test(n) ? title + "首尾不能出现下划线'_'" : /^\d+\d+\d$/.test(n) ? title + "不能全为数字" : void 0 : title + "不能含有特殊字符"
                                }, pass: [/^[\S]{6,12}$/, "密码必须6到12位，且不能出现空格"]
                            })
                        }, submitForm: function (n, e = null, t = null, a = !0) {
                            var i = [], s = [], c = n;
                            if (r.each(c, (function (n, e) {
                                if (console.log(n + ":" + e), /\[|\]|【|】/g.test(n)) {
                                    var t = n.match(/\[(.+?)\]/g);
                                    e = n.match("\\[(.+?)\\]")[1];
                                    var a = n.replace(t, "");
                                    r.inArray(a, i) < 0 && i.push(a), s[a] || (s[a] = []), s[a].push(e)
                                }
                            })), console.log(c), console.log(i), console.log(s), r.each(i, (function (n, e) {
                                var t = [];
                                r.each(s[e], (function (n, a) {
                                    t.push(a), delete c[e + "[" + a + "]"]
                                })), c[e] = t.join(",")
                            })), null == e) {
                                e = cUrl;
                                var u = r("form").attr("action");
                                o.isEmpty(u) ? null != n.id && (0 == n.id ? e += "/add" : n.id > 0 && (e += "/update")) : e = u
                            }
                            console.log(c);
                            var f = new FormData;
                            r.each(c, (function (n, e) {
                                f.append(n, e), console.log(n + "," + e)
                            })), console.log(f), l.ajaxPost(e, f, (function (n, e) {
                                if (e) return a && setTimeout((function () {
                                    var n = parent.layer.getFrameIndex(window.name);
                                    parent.layer.close(n)
                                }), 500), t && t(n, e), !1
                            }))
                        }, searchForm: function (n, e, t = "tableList") {
                            n.reload(t, {page: {curr: 1}, where: e.field})
                        }, initDate: function (n, e = null) {
                            if (Array.isArray(n)) for (var t in n) {
                                var i = n[t].split("|");
                                if (i[2]) var o = i[2].split(",");
                                var r = {};
                                if (r.elem = "#" + i[0], r.type = i[1], r.theme = "molv", r.range = "true" === i[3] || i[3], r.calendar = !0, r.show = !1, r.position = "absolute", r.trigger = "click", r.btns = ["clear", "now", "confirm"], r.mark = {
                                    "0-06-25": "生日",
                                    "0-12-31": "跨年"
                                }, r.ready = function (n) {
                                }, r.change = function (n, e, t) {
                                }, r.done = function (n, t, a) {
                                    e && e(n, t)
                                }, o) {
                                    var l = o[0];
                                    if (l) {
                                        var s = !isNaN(l);
                                        r.min = s ? parseInt(l) : l
                                    }
                                    var c = o[1];
                                    if (c) {
                                        var u = !isNaN(c);
                                        r.max = u ? parseInt(c) : c
                                    }
                                }
                                a.render(r)
                            }
                        }, showWin: function (n, e, t = 0, a = 0, i = [], o = 2, l = [], s = null, c = !1) {
                            var u = layui.layer.open({
                                title: n,
                                type: o,
                                area: [t + "px", a + "px"],
                                content: e,
                                shadeClose: c,
                                shade: .4,
                                skin: "layui-layer-admin",
                                success: function (n, e) {
                                    if (Array.isArray(i)) for (var t in i) {
                                        var a = i[t].split("=");
                                        layui.layer.getChildFrame("body", e).find("#" + a[0]).val(a[1])
                                    }
                                    s && s(e, 1)
                                },
                                end: function () {
                                    s(u, 2)
                                }
                            });
                            0 == t && (layui.layer.full(u), r(window).on("resize", (function () {
                                layui.layer.full(u)
                            })))
                        }, ajaxPost: function (n, e, a = null, i = "处理中,请稍后...") {
                            var o = null;
                            r.ajax({
                                type: "POST",
                                url: n,
                                data: e,
                                dataType: "json",
                                contentType: !1,
                                processData: !1,
                                beforeSend: function () {
                                    o = t.msg(i, {icon: 16, shade: .01, time: 0})
                                },
                                success: function (n) {
                                    if (0 != n.code) return t.close(o), t.msg(n.msg, {icon: 5}), !1;
                                    t.msg(n.msg, {icon: 1, time: 500}, (function () {
                                        t.close(o), a && a(n, !0)
                                    }))
                                },
                                error: function () {
                                    t.close(o), t.msg("AJAX请求异常"), a && a(null, !1)
                                }
                            })
                        }, ajaxGet: function (n, e, a = null, i = "处理中,请稍后...") {
                            var o = null;
                            r.ajax({
                                type: "GET",
                                url: n,
                                data: e,
                                contentType: "application/json",
                                dataType: "json",
                                beforeSend: function () {
                                    o = t.msg(i, {icon: 16, shade: .01, time: 0})
                                },
                                success: function (n) {
                                    if (0 != n.code) return t.msg(n.msg, {icon: 5}), !1;
                                    t.msg(n.msg, {icon: 1, time: 500}, (function () {
                                        t.close(o), a && a(n, !0)
                                    }))
                                },
                                error: function () {
                                    t.msg("AJAX请求异常"), a && a(null, !1)
                                }
                            })
                        }, formSwitch: function (n, t = "", a = null) {
                            e.on("switch(" + n + ")", (function (e) {
                                var i = this.checked ? "1" : "2";
                                o.isEmpty(t) && (t = cUrl + "/set" + n.substring(0, 1).toUpperCase() + n.substring(1));
                                var r = new FormData;
                                r.append("id", this.value), r.append(n.substring(0, 1).toLowerCase() + n.substring(1), i), console.log(r), l.ajaxPost(t, r, (function (n, e) {
                                    a && a(n, e)
                                }))
                            }))
                        }, uploadFile: function (n, e = null, a = "", r = "xls|xlsx", l = 10240, s = {}) {
                            o.isEmpty(a) && (a = cUrl + "/uploadFile"), i.render({
                                elem: "#" + n,
                                url: a,
                                auto: !1,
                                exts: r,
                                accept: "file",
                                size: l,
                                method: "post",
                                data: s,
                                before: function (n) {
                                    t.msg("上传并处理中。。。", {icon: 16, shade: .01, time: 0})
                                },
                                done: function (n) {
                                    return t.closeAll(), 0 == n.code ? t.alert(n.msg, {
                                        title: "上传反馈",
                                        skin: "layui-layer-molv",
                                        closeBtn: 1,
                                        anim: 0,
                                        btn: ["确定", "取消"],
                                        icon: 6,
                                        yes: function () {
                                            e && e(n, !0)
                                        },
                                        btn2: function () {
                                        }
                                    }) : t.msg(n.msg, {icon: 5}), !1
                                },
                                error: function () {
                                    return t.msg("数据请求异常")
                                }
                            })
                        }
                    };
                n("common", l)
            }))
        }
    }, e = {};
    !function t(a) {
        var i = e[a];
        if (void 0 !== i) return i.exports;
        var o = e[a] = {exports: {}};
        return n[a](o, o.exports, t), o.exports
    }(351)
})();