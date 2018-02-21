$(document).ready(function() {
    $("#form").submit(function(e) {
        e.preventDefault();

        var $error = $("#error");
        var $request = $("#request");
        var $result = $("#result");
        var $crt = $("#crt");

        $.ajax({
            type: "POST",
            url: "/api/request_certificate",
            data: JSON.stringify({
                csr: $("#csr").val()
            }),
            contentType: "application/json",
            dataType: "json",
            success: function(data) {
                $error.hide();
                $request.toggle();
                $result.toggle();
                $crt.val(data.certificate);

                $("#download").attr("href", "data:text/plain;charset=utf-8," + encodeURIComponent(data.certificate));
                $("#download").attr("download", "certificate.crt");
            },
            error: function(error) {
                $error.show();
            }
        });
    });

    $("#crt").click(function(e) {
        $(this).select();
    });
});
