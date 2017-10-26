var $clientsTable = $('#clientTable');
$(document).ready( function() {
	clientTable();
});

function clientTable() {
	var dt = $clientsTable.DataTable( {
        "processing": true,
		"serverSide": true,
		"keys": true,
		"autoWidth": false,
		"stateSave": false,
		"order": [[2, "asc"]],
        "ajax": "/client/table",
		"rowId": "id",
        "select": {
          "style":    'os',
          "selector": 'tr>td:nth-child(1), tr>td:nth-child(2), tr>td:nth-child(3), tr>td:nth-child(4)'
        },
		"columns": [
            {
				"data": null,
                "defaultContent": "",
				"width": "50",
				"orderable": false,
			},
            {
				"orderable": false,
                "data": "id",
				"defaultContent": "",
				"className": 'select-checkbox',
				"width": "90",
				'checkboxes': {
				    'selectRow': true
				}
			},
            { "name": "name", "data": "name" },
            { "name": "villa", "data": "villa" },
			{
				"data": null,
                "defaultContent": '<i class="icon-trash deleteActivityType"></i>',
				"width": '50',
				"orderable": false,
			}
		],
		"dom": '<"datatable-header"fl><"datatable-scroll-wrap"t><"datatable-footer"ip>',
        "language": {
            search: '<span>Filter:</span> _INPUT_',
            searchPlaceholder: 'Type to filter...',
            lengthMenu: '<span>Show:</span> _MENU_',
            paginate: { 'first': 'First', 'last': 'Last', 'next': '&rarr;', 'previous': '&larr;' }
        },
		"drawCallback": function () {
			dt.column(0, {search:'applied', order:'applied'}).nodes().each( function (cell, i) {
				var start = dt.page.info().start;
				cell.innerHTML = start+i+1;
			});
		}
    });

    dt.on( 'select deselect draw', function ( e, dt, type, indexes ) {
		console.log($("tbody .dt-checkboxes:checked").length);
		if ( $("tbody .dt-checkboxes:checked").length == $('tbody .dt-checkboxes').length) {
			$("thead tr").removeClass("indeterminate");
			$("thead tr").addClass("selected");
		} else {
			$("thead tr").removeClass("selected");
			if($("tbody .dt-checkboxes:checked").length > 0) {
				$("thead tr input").prop({
					indeterminate: true,
					checked: false
				});
				$("thead tr").addClass("indeterminate");
			} else {
				$("thead tr").removeClass("selected");
				$("thead tr").removeClass("indeterminate");
			}
		}

	});


	$('.activityTypeNewBtn').on('click', function() {
		addNewActivity(dt2);
	});

	$(".activityType").on('keydown', function(event) {
		if (event.keyCode == 13) {
			addNewActivity(dt2);
		}
	})


	$('#activityTypeTable').on('click', '.deleteActivityType', function() {
		var activityTypeID = $(this).closest('tr').attr("id");
		$.ajax({
			type:"DELETE",
			url:"/activitytype/"+activityTypeID,
			success: function(data) {
				if(data.success) {
					dt2.ajax.reload(null, false);
					new PNotify({
						title: 'Value Deleted',
						text: 'Value was deteled successfully!',
						type: 'success',
						icon: 'icon-checkmark3'
					});
				} else {
					new PNotify({
						title: 'Value was not deleted!',
						text: 'There was an error deleting the value, please try again.',
						type: 'error',
						icon: 'icon-blocked'
					});
				}

			},
			dataType: 'json',
		});
	});

    $('.newClient').on('click', function() {
        newClientHandler();
    });

    $('.popupContainer').on('click', '.addClient', function() {
        var formSelector = $(".editBusinessForm");
    	var b = {};
    	$.each($(formSelector).serializeArray(), function(i, field) {
    		b[field.name] = field.value;
    	});

        $.ajax({
		type:"POST",
		url:"/client/add",
			data: b,
			success: function(data) {
                if(data.status === true) {
                    $(".clientModal").modal("hide");
                    dt.ajax.reload(null, false);
                }

			},
			dataType: 'json',
		  });
    });

    $(".popupContainer").on('hidden.bs.modal', '.modal', function (e) {

	})
}

function newClientHandler() {
    $('.popupContainer').html(businessEditModal);
    $('.clientModal').modal();
}

var businessEditModal = '<div class="modal fade clientModal ">'+
						'<div class="modal-dialog modal-lg">'+
							'<div class="modal-content">'+
								'<div class="modal-header">'+
									'<button type="button" class="close" data-dismiss="modal">×</button>'+
									'<h5 id="businessEditModalLabel" class="modal-title">Update business - <strong></strong></h5>'+
								'</div>'+

								'<div class="modal-body">'+
									'<div class="row">'+
									'<div class="col-sm-12">'+
									'<ul class="nav nav-tabs">'+
										'<li class="active"><a href="#nameAddressStepTab" data-toggle="tab">Name & Address</a></li>'+
										'<li><a href="#detailsTab" data-toggle="tab">Details</a></li>'+
										'<li><a href="#description" data-toggle="tab">Description</a></li>'+
										'<li><a href="#social" data-toggle="tab">Social</a></li>'+
									'</ul>'+
									'<form class="editBusinessForm">' +
										'<div class="tab-content">'+
											'<div class="tab-pane active" id="nameAddressStepTab">'+
												'<div class="form-group">'+
													'<div class="row">'+
														'<div class="col-sm-6">'+
															'<label>Name</label>'+
															'<input type="text" placeholder="" name="name" class="form-control clientName">'+
														'</div>'+

														'<div class="col-sm-6">'+
															'<label>No.</label>'+
															'<input type="text" placeholder="" name="no" class="form-control clientNo">'+
														'</div>'+
													'</div>'+
												'</div>'+

												'<div class="form-group">'+
													'<div class="row">'+
														'<div class="col-sm-6">'+
															'<label>Villa</label>'+
															'<input type="text" placeholder="" name="villa" class="form-control clientVilla">'+
														'</div>'+

														'<div class="col-sm-6">'+

														'</div>'+
													'</div>'+
												'</div>'+
											'</div>'+

											'<div class="tab-pane" id="detailsTab">'+
                                            'tab2'+
											'</div>'+

											'<div class="tab-pane" id="description">'+
                                                'tab 3'+
											'</div>'+

											'<div class="tab-pane" id="social">'+
                                                'tab 4'+
											'</div>'+
										'</div>'+
									'</form>'+
									'</div>'+
									'</div>'+
								'</div>'+
								'<div class="modal-footer">'+
									'<button type="button" class="btn btn-link" data-dismiss="modal">Close</button>'+
									'<button type="submit" class="btn btn-primary addClient">Add</button>'+
								'</form>'+
							'</div>'+
						'</div>'+
					'</div>'+
				'</div>';
