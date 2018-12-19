import scala.io.Source
import java.io.PrintWriter
import java.io.File
import scala.collection.mutable
val inputDirPath = "tag_result/"
val crossNum = 3
val mode = 2//mode 1代表以人为单位统计准确，mode2代表以包为单位
val fileList = (new File(inputDirPath)).listFiles.filter(_.getName.endsWith(".tagged")).groupBy(_.getName.split("_").take(mode).mkString("_"))
val personResultMap = fileList.mapValues{arr=>var startIndex = 0;var urlIdMap = mutable.HashMap[String,Int]();arr.foreach{file=>val lines = Source.fromFile(file).getLines.toArray;lines.foreach(l=>urlIdMap.put(l.split("\t")(1),startIndex + l.split("\t")(0).toInt));startIndex += lines.size};urlIdMap}
val canScoreData = personResultMap.toList.flatMap(e=>e._2.toList.map(f=>(f._1,(e._1,f)))).groupBy(_._1).filter(_._2.size >= crossNum).toList.flatMap(_._2.map(_._2)).groupBy(_._1).mapValues(_.map(_._2).toMap)
val personMergePairMap = canScoreData.map(e=>(e._1,e._2.toList.map(_.swap).groupBy(_._1).filter(_._2.size >=2).flatMap(_._2.map(_._2).combinations(2).map(_.toSet)).toSet))
val allShouldMergePairMap = personMergePairMap.toList.map(_._2).combinations(crossNum/2 + 1).map(_.reduce(_ & _)).reduce(_ | _).flatMap(e=>e.map(f=>(f,e))).groupBy(_._1).mapValues(_.map(_._2)) //所有应该和的
val personShouldMergePairMap = canScoreData.mapValues(e=>(allShouldMergePairMap filterKeys e.keySet).values.reduce(_ | _)) //一个人所有标注里面该和的。
//该合的合了
val TPs = personShouldMergePairMap.map(e=>(e._1,e._2 & personMergePairMap.getOrElse(e._1,Set()))).mapValues(_.size)
//不该合的合了
val FPs = personMergePairMap.map(e=>(e._1,e._2 -- personShouldMergePairMap.getOrElse(e._1,Set()))).mapValues(_.size)
//该合的没合
val FNs = personShouldMergePairMap.map(e=>(e._1,e._2 -- personMergePairMap.getOrElse(e._1,Set()))).mapValues(_.size)
//准确率
val precisions = TPs.map(e=>(e._1,e._2.toDouble/(e._2+FPs.getOrElse(e._1,0))))
//召回率
val recalls = TPs.map(e=>(e._1,e._2.toDouble/(e._2+FNs.getOrElse(e._1,0))))
val n = 1.0
//F值
val FNmeasure = precisions.map(e=>(e._1,(1+n*n)*e._2*recalls.getOrElse(e._1,0.0)/(n*n*(e._2 + recalls.getOrElse(e._1,0.0)))))
val outFile = new PrintWriter("best_tag.txt")
val resultPair = personShouldMergePairMap.map(_._2).reduce(_|_)
