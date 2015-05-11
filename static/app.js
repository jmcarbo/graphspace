
$.fn.serializeObject = function()
{
    var o = {};
    var a = this.serializeArray();
    $.each(a, function() {
        if (o[this.name] !== undefined) {
            if (!o[this.name].push) {
                o[this.name] = [o[this.name]];
            }
            o[this.name].push(this.value || '');
        } else {
            o[this.name] = this.value || '';
        }
    });
    return o;
};


function getUrlParameter(sParam)
{
    var sPageURL = window.location.search.substring(1);
    var sURLVariables = sPageURL.split('&');
    for (var i = 0; i < sURLVariables.length; i++) 
    {
        var sParameterName = sURLVariables[i].split('=');
        if (sParameterName[0] == sParam) 
        {
            return sParameterName[1];
        }
    }
}         

function updateImage(data) {
   var img_url = '/image/graph.' + data.output + '?id=' + data.id
   var link_text = '<a href="?id=' + data.id + '">share<a/>'
   link_text += '&nbsp; <a href="' + img_url + '">image<a/>'
   $('#link').html(link_text)

   $('#graph').html('<a href="' + img_url + '"><img src="data:' + data.content_type + ';base64,' + data.image + '"/></a>')
   $('#text').val(data.text)
   $('#description').val(data.description)


   if (data.format) {
      $("#form input[name='format'][value=" + data.format + "]").attr('checked', 'checked')
   }
   if (data.output) {
      $("#form input[name='output'][value=" + data.output + "]").attr('checked', 'checked')
   }
   refreshRecent()
}

function deleteImage(id) {
   $.ajax({
      type: "POST",
      url: "/api/delete?id=" + id,
      dataType: "json",
      success: function(data) {
         $('#img-' + id).hide()
      }
   })
}

function refreshRecent() {
   $.ajax({
      type: "GET",
      url: "/api/list",
      dataType: "json",
      success: function(data) {
         var buf = ''
         $.each(data, function() {
            if (this.description == '') {
               this.description = '<i>no description</i>'
            }
            var delete_cmd = 'javascript:deleteImage("' + this.id + '");'
            var v = "<div id='img-" + this.id + "' class='col' style='text-align: center;'>"
            v += "<a href='?id=" + this.id + "'>" + this.description  + "</a>"
            v += " &nbsp; <a href='" + delete_cmd + "'>X</a><br>"

            v += "<a href='?id=" + this.id + "'>"
            v += "<img width=150 height=150 src='/image/graph.png?id=" + this.id + "'/></a>"
            v += "</div>"
            buf += v
         })
         $('#recent').html(buf)
      }
   })
}