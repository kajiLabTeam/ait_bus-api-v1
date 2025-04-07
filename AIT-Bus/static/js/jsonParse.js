function onDisplay(hour, hour2, hour3, minutes1, minutes2, id1, id2, id3, isFirst) {
    /*
    出発時間を表示する関数
    :param hour:        # int：現在時間(時)
    :param hour2:       # int：現在時間(時) or 現在時間+1(時)
    :param hour3:       # int：現在時間(時) or 現在時間+1(時) or 現在時間+2(時)
    :param minutes1:    # int：次の出発時間(分)
    :param minutes2:    # int：次の次の出発時間(分)
    :param id1:         # str：表示する部分のid
    :param id2:         # str：表示する部分のid
    :param id3:         # str：表示する部分のid
    :param isFirst:     # bool：最初のバスかどうか
    */

    // 次の出発時刻を表示する
    if (minutes1 != -1 && !isFirst) {
        document.getElementById(id1).innerHTML = hour2 + ":" + ('00' + minutes1).slice(-2);
    } else {
        if (hour <= 7) {
            document.getElementById(id1).innerHTML = "[始発] 0" + hour3 + ":" + ('00' + minutes1).slice(-2);
        } else {
            document.getElementById(id1).innerHTML = "本日の運行は終了しました";
        }
    }

    // 次の次の出発時間を表示する
    if (minutes2 != -1 && hour > 7) {
        if (document.getElementById(id2).style.display == "none"
            || document.getElementById(id3).style.display == "none") {
            document.getElementById(id2).style.display = "";
            document.getElementById(id3).style.display = "";
        }
        document.getElementById(id2).innerHTML = "after the next";
        document.getElementById(id3).innerHTML = hour3 + ":" + ('00' + minutes2).slice(-2);
    } else {
        document.getElementById(id2).style.display = "none";
        document.getElementById(id3).style.display = "none";
    }
}

function reload() {
    var nowTime = new Date();
    var nowSec = nowTime.getSeconds();
    if (nowSec == 1) {
        location.reload()
    }
}

//時刻表の初期表示は非表示
document.getElementById("A1").style.display = "none";
document.getElementById("B1").style.display = "none";
document.getElementById("C1").style.display = "none";
document.getElementById("A2").style.display = "none";
document.getElementById("B2").style.display = "none";
document.getElementById("C2").style.display = "none";

async function main() {
    // 現在の時間(hour)を取得
    const date = new Date();
    const hour = date.getHours();

    // 今日の日付を取得
    const nextBusInfoUrl = "https://bus-api.bigbell.dev/api/v1/nextbus?offset=0";
    const afterNextBusInfoUrl = "https://bus-api.bigbell.dev/api/v1/nextbus?offset=1";
    const nextBusInfo = await fetch(nextBusInfoUrl).then(response => response.json());
    const afterNextBusInfo = await fetch(afterNextBusInfoUrl).then(response => response.json());
    console.info(nextBusInfo);
    console.info(afterNextBusInfo);
    console.log("================================");

    // もし今日のバスがなければページを遷移する
    if (nextBusInfo["schedule"] == "") {
        // window.location.href = "./nobus.html?ver=2023_2";
        document.getElementById("to_daigaku").innerHTML = "本日バスの運行はありません";
        document.getElementById("to_yakusa").innerHTML = "本日バスの運行はありません";
        // document.getElementById("to_daigaku").innerHTML = "本年度の<br>運行スケジュールが<br>未発表なため不明です";
        // document.getElementById("to_yakusa").innerHTML = "本年度の<br>運行スケジュールが<br>未発表なため不明です";
        document.getElementById("after-the-next-daigaku").style.display = "none";
        document.getElementById("next_to_daigaku").style.display = "none";
        document.getElementById("after-the-next-yakusa").style.display = "none";
        document.getElementById("next_to_yakusa").style.display = "none";
        document.getElementById("first").style.display = "none";
        document.getElementById("second").style.display = "none";
        document.getElementById("third").style.display = "none";
        document.getElementById("navSchedule").style.display = "none";
        document.getElementById("navTableToAIT").style.display = "none";
        document.getElementById("navTableToYakusa").style.display = "none";
        return;
    }

    let schedule = "";
    let nextHourToAIT = 0;
    let nextHourToYakusa = 0;
    let nextMinuteToAIT = 0;
    let nextMinuteToYakusa = 0;
    let afterNextHourToAIT = 0;
    let afterNextHourToYakusa = 0;
    let afterNextMinuteToAIT = 0;
    let afterNextMinuteToYakusa = 0;
    if (nextBusInfo["busState"]["IsExist"] || nextBusInfo["busState"]["IsFirst"]) {
        schedule = nextBusInfo["schedule"];
        nextHourToAIT = nextBusInfo["nextHourToAIT"];
        nextHourToYakusa = nextBusInfo["nextHourToYakusa"];
        nextMinuteToAIT = nextBusInfo["nextMinuteToAIT"];
        nextMinuteToYakusa = nextBusInfo["nextMinuteToYakusa"];
    } else {
        nextMinuteToAIT = -1;
        nextMinuteToYakusa = -1;
        schedule = nextBusInfo["schedule"];
    }

    if (afterNextBusInfo["busState"]["IsExist"] || afterNextBusInfo["busState"]["IsFirst"]) {
        schedule = afterNextBusInfo["schedule"];
        afterNextHourToAIT = afterNextBusInfo["nextHourToAIT"];
        afterNextHourToYakusa = afterNextBusInfo["nextHourToYakusa"];
        afterNextMinuteToAIT = afterNextBusInfo["nextMinuteToAIT"];
        afterNextMinuteToYakusa = afterNextBusInfo["nextMinuteToYakusa"];
    } else {
        afterNextMinuteToAIT = -1;
        afterNextMinuteToYakusa = -1;
    }

    if (afterNextBusInfo["busState"]["IsFirst"]) {
        schedule = afterNextBusInfo["schedule"];
        afterNextHourToAIT = afterNextBusInfo["nextHourToAIT"];
        afterNextHourToYakusa = afterNextBusInfo["nextHourToYakusa"];
        afterNextMinuteToAIT = afterNextBusInfo["nextMinuteToAIT"];
        afterNextMinuteToYakusa = afterNextBusInfo["nextMinuteToYakusa"];
    }

    // ダイヤを表示する
    if (schedule == "T2") {
        document.getElementById("daiya").innerHTML = "今日は特②ダイヤです";
    } else {
        document.getElementById("daiya").innerHTML = "今日は" + schedule + "ダイヤです";
    }

    // 出発時刻を表示する(大学行き)
    onDisplay(hour, nextHourToAIT, afterNextHourToAIT, nextMinuteToAIT, afterNextMinuteToAIT, "to_daigaku", "after-the-next-daigaku", "next_to_daigaku", nextBusInfo["busState"]["IsFirst"]);

    // 出発時刻を表示する(八草行き)
    onDisplay(hour, nextHourToYakusa, afterNextHourToYakusa, nextMinuteToYakusa, afterNextMinuteToYakusa, "to_yakusa", "after-the-next-yakusa", "next_to_yakusa", nextBusInfo["busState"]["IsFirst"]);

    // その日のダイヤの時刻表を表示
    const A1 = document.getElementById("A1");
    const B1 = document.getElementById("B1");
    const C1 = document.getElementById("C1");
    // const T2_toDaigaku = document.getElementById("T2_toDaigaku");
    const A2 = document.getElementById("A2");
    const B2 = document.getElementById("B2");
    const C2 = document.getElementById("C2");
    // const T2_toYakusa = document.getElementById("T2_toYakusa");

    if (schedule == "A") {
        A1.style.display = "block";
        A2.style.display = "block";
    } else if (schedule == "B") {
        B1.style.display = "block";
        B2.style.display = "block";
    } else if (schedule == "C") {
        C1.style.display = "block";
        C2.style.display = "block";
    } else if (schedule == "T2") {
        // T2_toDaigaku.style.display = "block";
        // T2_toYakusa.style.display = "block";
    }

    // 画像のリンクを生成
    // document.getElementById("to_daigaku_image").innerHTML = '<span class="image"><a href="https://ogane.sakura.ne.jp/AIT-BusInfo/static/images/' + daiya + '1.jpg"><img src="static/images/' + daiya + '1.jpg" alt=""></a></span><br></br>';
    // document.getElementById("to_yakusa_image").innerHTML = '<span class="image"><a href="https://ogane.sakura.ne.jp/AIT-BusInfo/static/images/' + daiya + '2.jpg"><img src="static/images/' + daiya + '2.jpg" alt=""></a></span><br></br>';
}

// reload once when the new Service Worker starts activating
var refreshing;
navigator.serviceWorker.addEventListener('controllerchange',
    function () {
        if (refreshing) return;
        refreshing = true;
        console.log('Service Workerをアップデートしたのでリロードします!!');
        window.location.reload();
    }
);

// setInterval('reload()', 1000);
main();

// 毎分、現在の秒数が0の時にmain()を実行する
setInterval(function () {
    var nowTime = new Date();
    var nowSec = nowTime.getSeconds();
    if (nowSec == 0) {
        main();
    }
}, 1000);