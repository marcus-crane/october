interface go {
  "main": {
    "KoboService": {
		DetectKobos():Promise<Array<Kobo>>
		ListDeviceBookmarks():Promise<Error>
		ListDeviceContent():Promise<Error>
		OpenDBConnection(arg1:string):Promise<Error>
		PromptForLocalDBPath():Promise<Error>
		SelectKobo(arg1:string):Promise<boolean>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
