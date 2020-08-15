$(function() {
	layui.use(['form','element']),
		function() {
			layer = layui.layer;
			element = layui.element;
		}
	function xxxclose() {
		var index = parent.layer.getFrameIndex(window.name)
		parent.layer.close(index);
	}
});

function xxclose() {
	var index = parent.layer.getFrameIndex(window.name)
	parent.layer.close(index);
}


/*add.html*/
function add() {
	layer.open({
		type:2,
		area:['800px','800px'],
		fixed: false,
		shadeClose: true,
		shade: 0.4,
		content: "/admin/add",
		title: "Student-Admin"
	});
}

function submit() {
	var uid = $("#id").val();
	var uname = $("#name").val();
	var umajor = $("#major").val();
	var usocre = $("#socre").val();
	var ubirthday = $("#birthday").val();
	var usex = $("#sex").val();
	$.ajax({
		url: "/admin/add",
		dataType: "json",
		type: "post",
		data: {
			id: uid,
			name: uname,
			major: umajor,
			socre: usocre,
			birthday: ubirthday,
			sex: usex
		},
		success: function() {
			alert("添加成功");
			location.reload();
		},
		error: function() {
			alert("添加失败");
		}
    });
}

/*index.html*/
function nextpage() {
	var page = Number($("#next").val());
	var pages = Number($("#pages").val()) - 1;
	var url = "/admin/index?page=";
	if (page<pages) {
		page +=1;
		url = url + page;
		location.href = url;
	} else {
		$("#next").attr("disabled", ture);
	}
}

function prevpage() {
	var page = Number($("#prev").val());
	var pages = Number($("#pages").val());
	var url = "/admin/index?page=";		
	if (page<1 || pages == 1) {
		$("#prev").attr("disabled", ture);
	} else {
		page -=1;
		url = url + page;
		location.href = url;
	}
}

function deletes() {
	var r = confirm("是否删除当前学生信息");
	var id = $("#del").val();
	if (r == true)
	{
		deletes1(id)
	}
}

function deletes1(uid) {
	$.ajax({
		url: "/admin/delete",
		type: "post",
		data: {
			delid: uid,
		},
		success: function() {
			alert("删除成功");
			location.href = "/admin/index";
		},
		error: function() {
			alert("删除失败");
		}
	});
}


/*preview.html*/
function edit(id) {
	layer.open({
		type:2,
		area:['800px', '800px'],
		shadeClose: true,
		shade: 0.4,
		content: "/admin/edit?id=" + id,
		title: id
	});
}

function edits(id) {
	var uname = $("#name").val();
	var umajor = $("#major").val();
	var usocre = $("#socre").val();
	var ubirthday = $("#birthday").val();
	var usex = $("#sex").val();
	$.ajax({
		url: "/admin/edit",
		dataType: "json",
		type: "post",
		data: {
			id: id,
			name: uname,
			major: umajor,
			socre: usocre,
			birthday: ubirthday,
			sex: usex
		},
		success: function() {
			xclose();
			alert("修改成功");
			window.parent.location.reload();
		},
		error: function() {
			alert("添加失败");
		}
	});
}
function xclose() {
	var index = parent.layer.getFrameIndex(window.name)
	parent.layer.close(index);
}
