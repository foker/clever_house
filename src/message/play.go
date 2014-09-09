package message

import (
	"crypto/md5"
	"os"
	"os/exec"
	"time"
)

func Play(messages []string, backgroundMusicFile string){
	var mFiles []string

	for i := 0; i < len(messages); i++ {
		messageFile := getMD5(messages[i]) + ".mp3"
		mFiles = append(mFiles, messageFile)
		//Скачиваем файл с сообщением
		if err := exec.Command("wget", "-q", "-U Mozilla", "-O"+messageFile, "http://translate.google.com/translate_tts?ie=UTF-8&tl=ru&q="+messages[i]).Run(); err != nil {
			log.Println(messages[i])
			return err
		}
	}

	//Запускаем фоновую музыку
	backgroundSong := exec.Command("mplayer", backgroundMusicFile)

	wr, _ := backgroundSong.StdinPipe()

	if err := backgroundSong.Start(); err != nil {
		return err
	}

	//Ждем когда чучуть поиграет
	time.Sleep(1000 * time.Millisecond)

	//Убавляем громкость
	for i := 0; i < 5; i++ {
		time.Sleep(300 * time.Millisecond)
		wr.Write([]byte("/"))
	}

	for i := 0; i < len(mFiles); i++ {
		//Проигрываем сообщение
		if err := exec.Command("mplayer", mFiles[i]).Run(); err != nil {
			return err
		}
		//Удаляем файл сообщения
		if err := os.Remove(mFiles[i]); err != nil {
			return err
		}
	}

	//Прибавляем громкость обратно
	for i := 0; i < 5; i++ {
		time.Sleep(300 * time.Millisecond)
		wr.Write([]byte("*"))
	}

	//Слушаем еще немного музыки
	time.Sleep(3000 * time.Millisecond)

	//Убавляем громкость
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		wr.Write([]byte("/"))
	}
	//Завершаем воспроизведение
	wr.Write([]byte("q"))

	if err := backgroundSong.Wait(); err != nil {
		return err
	}

	return nil


}
