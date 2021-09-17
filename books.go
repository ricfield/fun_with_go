package main
import ("fmt"; "strconv")

func main() {
   var publisher string
   var writer string
   var artist string
   var title string
   var year int
   var pageNumber int
   var grade float32

   title = "Mr.GoToSleep"
   writer = "Tracey Hatchet"
   artist = "Jewel Tampson"
   publisher = "DizzyBooks Publishing Inc."
   year = 1997
   pageNumber = 14
   grade = 6.5

   fmt.Println("Publisher: " + publisher +
              "\nWriter: " + writer +
              "\nArtist: " + artist +
              "\nTitle: " + title +
              "\nYear: " + strconv.Itoa(year) +
              "\nTotal Pages: " + strconv.Itoa(pageNumber) +
              "\nGrade:", grade)

   writer = "Ryan N. Shawn"
   artist = "Phoebe Paperclips"
   year = 2013
   pageNumber = 160
   grade = 9.0

   fmt.Println("Publisher: " + publisher +
              "\nWriter: " + writer +
              "\nArtist: " + artist +
              "\nTitle: " + title +
              "\nYear: " + strconv.Itoa(year) +
              "\nTotal Pages: " + strconv.Itoa(pageNumber) +
              "\nGrade:", grade)

}
