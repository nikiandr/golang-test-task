function addField (argument) {
    

    input_name = document.getElementById("name").value;
    input_email = document.getElementById("email").value;
    var email_re = /^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/
    if (!email_re.test(input_email)) {
        alert("Wrong email format.")
    }
    else if (input_name != '' && input_email != ''){
        var members_table = document.getElementById("members");

        var currentIndex = members_table.rows.length;

        var newRow = members_table.insertRow(-1);

        var today = new Date();

        var index_cell = newRow.insertCell(0);
        var name_cell = newRow.insertCell(1);
        var email_cell = newRow.insertCell(2);
        var registration_date_cell = newRow.insertCell(3);

        index_cell.innerHTML = currentIndex;
        name_cell.innerHTML = input_name;
        email_cell.innerHTML = input_email;
        registration_date_cell.innerHTML = today.toISOString().split('T')[0];
    }

    document.getElementById("name").value = '';
    document.getElementById("email").value = '';
}