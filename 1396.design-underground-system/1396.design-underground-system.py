#
# @lc app=leetcode id=1396 lang=python3
#
# [1396] Design Underground System
#

# @lc code=start


class UndergroundSystem:

    def __init__(self):
        self.starts = {}
        self.ends = {}

    def checkIn(self, id: int, stationName: str, t: int) -> None:
        if self.starts.__contains__(stationName):
            if self.starts[stationName].__contains__(id):
                self.starts[stationName][id].append(t)
            else:
                self.starts[stationName][id] = [t]
        else:
            self.starts[stationName] = {}
            self.starts[stationName][id] = [t]

    def checkOut(self, id: int, stationName: str, t: int) -> None:
        if self.ends.__contains__(stationName):
            if self.ends[stationName].__contains__(id):
                self.ends[stationName][id].append(t)
            else:
                self.ends[stationName][id] = [t]
        else:
            self.ends[stationName] = {}
            self.ends[stationName][id] = [t]

    def getAverageTime(self, startStation: str, endStation: str) -> float:
        if not self.starts.__contains__(startStation) or not self.ends.__contains__(endStation):
            return 0
        times, count = 0, 0
        starts, ends = self.starts[startStation], self.ends[endStation]
        for k in starts:
            if not ends.__contains__(k):
                continue
            endTimes = ends[k][:]
            for st in starts[k]:
                if len(endTimes) > 0:
                    times += (endTimes[0]-st)
                    count += 1
                    del endTimes[0]
        if times == 0:
            return 0
        return times/count


# Your UndergroundSystem object will be instantiated and called as such:
# obj = UndergroundSystem()
# obj.checkIn(id,stationName,t)
# obj.checkOut(id,stationName,t)
# param_3 = obj.getAverageTime(startStation,endStation)
