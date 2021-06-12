function GetOldResult() {
  return document.getElementById("old_result").innerText;
}

function PrintOldResult(so) {
  document.getElementById("old_result").innerText = so;
}

function GetResult() {
  return document.getElementById("result").innerText;
}

function PrintResult(so) {
  document.getElementById("result").innerText = so;
}

var operator = document.getElementsByClassName("operator");
for (var i = 0; i < operator.length; i++) {
  operator[i].addEventListener("click", function () {
    if (this.id == "clearAll") {
      PrintResult("");
      PrintOldResult("");
    } else if (this.id == "clear") {
      let result = GetResult().toString();
      if (result) {
        result = result.substr(0, result.length - 1);
        PrintResult(result);
      }
    } else {
      var result = GetResult();
      var old_result = GetOldResult();
      if (result != "") {
        old_result = old_result + result;
        if (this.id == "=") {
          var ket_qua_cuoi = eval(old_result);
          PrintResult(ket_qua_cuoi);
          PrintOldResult("");
        } else {
          old_result = old_result + this.id;
          PrintOldResult(old_result);
          PrintResult("");
        }
      }
    }
  });
}

var number = document.getElementsByClassName("number");
for (var i = 0; i < number.length; i++) {
  number[i].addEventListener("click", function () {
    var result = GetResult();
    if (result != NaN) {
      result = result + this.id;
      PrintResult(result);
    }
  });
}
